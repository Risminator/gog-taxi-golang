package main

import (
	"context"

	"github.com/Risminator/gog-taxi-golang/internal/controllers"
)

func main() {
	ctx := context.Background()
	ch := make(chan int)
	controllers.CreateServer(ctx, ch)
	<-ch
}
