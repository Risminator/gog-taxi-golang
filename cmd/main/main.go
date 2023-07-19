package main

import (
	"context"

	"github.com/Risminator/gog-taxi-golang/internal/app"
)

func main() {
	ctx := context.Background()
	ch := make(chan int)
	app.CreateServer(ctx, ch)
	<-ch
}
