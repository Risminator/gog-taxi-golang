package controllers

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

	"github.com/Risminator/gog-taxi-golang/internal/app"
	"github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

const (
	httpPort = ":18080"
)

func NewHTTPServer(port string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	api := handler.Group("/api/v1")
	httpgin.AppRouter(api, a)
	s := &http.Server{Addr: port, Handler: handler}
	return s
}

func CreateServer(ctx context.Context, ch chan int) *http.Server {
	// ToDo: add repositories
	a := app.NewApp()
	return CreateServerWithExternalApp(ctx, ch, a)
}

func CreateServerWithExternalApp(ctx context.Context, ch chan int, a app.App) *http.Server {
	httpServer := NewHTTPServer(httpPort, a)
	eg, ctx := errgroup.WithContext(ctx)

	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		eg.Go(func() error {
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
			defer log.Printf("close http server listening on %s\n", httpServer.Addr)

			errCh := make(chan error)

			defer func() {
				shCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
				defer cancel()

				if err := httpServer.Shutdown(shCtx); err != nil {
					log.Printf("can't close http server listening on %s: %s", httpServer.Addr, err.Error())
				}

				close(errCh)
			}()

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

		if err := eg.Wait(); err != nil {
			log.Printf("gracefully shutting down the servers: %s\n", err.Error())
		}

		log.Println("servers were successfully shutdown")

		ch <- 0
	}()
	time.Sleep(time.Millisecond * 30)
	return httpServer
}
