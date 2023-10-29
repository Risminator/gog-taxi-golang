package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	v0 "github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin/v0"
	v1 "github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin/v1"
	"github.com/Risminator/gog-taxi-golang/internal/infrastructure/datastore"
	"github.com/Risminator/gog-taxi-golang/internal/infrastructure/repository"
	"github.com/Risminator/gog-taxi-golang/internal/infrastructure/webapi"
	"github.com/Risminator/gog-taxi-golang/internal/infrastructure/websockets"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

const (
	httpPort = ":18080"
)

func NewHTTPServer(port string, hu usecase.Hello, cu usecase.Customer, du usecase.Dock,
	dru usecase.Driver, vu usecase.Vessel, ru usecase.TaxiRequest,
	rws v1.TaxiRequestWsGateway, routeUsecase usecase.Route) *http.Server {
	gin.SetMode(gin.ReleaseMode)

	// Initialize handler with logger and recovery
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Initialize api routes
	// in future we won't use several api versions, it's just for testing purposes
	api := handler.Group("/api")
	{
		v0.NewRouter(api, hu)
		v1.NewRouter(api, cu, du, dru, vu, ru, rws, routeUsecase)
	}

	s := &http.Server{Addr: port, Handler: handler}
	return s
}

func CreateServer(ctx context.Context, ch chan int) *http.Server {
	db := datastore.NewDB()

	hu := usecase.NewHelloUsecase()

	cr := repository.NewCustomerRepository(db)
	cu := usecase.NewCustomerUsecase(cr)

	dr := repository.NewDockRepository(db)
	du := usecase.NewDockUsecase(dr)

	drr := repository.NewDriverRepository(db)
	dru := usecase.NewDriverUsecase(drr)

	vr := repository.NewVesselRepository(db)
	vu := usecase.NewVesselUsecase(vr)

	rbWebApi := webapi.NewRouterBrouter()
	routeUsecase := usecase.NewRouteUsecase(rbWebApi)

	rr := repository.NewTaxiRequestRepository(db)
	ru := usecase.NewTaxiRequestUsecase(rr, rbWebApi)

	wsManager := websockets.NewManager(ctx)
	rws := websockets.NewWsTaxiRequestHandler(wsManager, ru, vu)

	httpServer := NewHTTPServer(httpPort, hu, cu, du, dru, vu, ru, rws, routeUsecase)
	eg, ctx := errgroup.WithContext(ctx)

	// Signals capture and graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		eg.Go(func() error {
			// Signal capture
			select {
			case s := <-sigQuit:
				log.Printf("captured signal: %v\n", s)
				return fmt.Errorf("captured signal: %v", s)
			case <-ctx.Done():
				return nil
			}
		})

		eg.Go(func() error {
			log.Printf("starting http server, listening on %s\n", httpServer.Addr)
			defer log.Printf("closing http server listening on %s\n", httpServer.Addr)

			errCh := make(chan error)

			// Trying to shutdown the server at the end of the function
			defer func() {
				shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()

				if err := httpServer.Shutdown(shCtx); err != nil {
					log.Printf("can't close http server listening on %s: %s", httpServer.Addr, err.Error())
				}

				close(errCh)
			}()

			// Trying to shutdown the db at the end of the function
			defer func() {
				instance, _ := db.DB()
				instance.Close()
				log.Printf("database closed")
			}()

			// Listening on a separate goroutine
			go func() {
				if err := httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
					errCh <- err
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errCh:
				return fmt.Errorf("http server can't listen and serve requests: %w", err)
			}
		})

		// Graceful shutdown
		if err := eg.Wait(); err != nil {
			log.Printf("gracefully shutting down the servers: %s\n", err.Error())
		}

		log.Println("servers were successfully shutdown")

		ch <- 0
	}()
	time.Sleep(time.Millisecond * 30)
	return httpServer
}
