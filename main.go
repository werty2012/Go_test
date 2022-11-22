package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	done := make(chan struct{})
	s := Server{}
	s.Run()

	go func() {
		var s string
		fmt.Scanln(&s)
		close(done)
	}()

	<-done
	s.Shutdown(ctx)
}
