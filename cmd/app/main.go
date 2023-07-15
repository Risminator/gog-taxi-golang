package main

import (
	"context"

	"github.com/Risminator/gog-taxi-golang/internal/controllers/httpgin"
)

func main() {
	ctx := context.Background()
	ch := make(chan int)
	httpgin.CreateServer(ctx, ch)
	<-ch
}
