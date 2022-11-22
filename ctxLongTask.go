package main

import (
	"context"
	"fmt"
	"time"
)

func run() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()

	processLongTask(ctx)

}

func processLongTask(ctx context.Context) {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("done processing")
	case <-ctx.Done():
		fmt.Println("cancelled")

	}

}
