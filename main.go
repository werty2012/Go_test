package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()

	done := make(chan struct{})

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	s := Server{}
	s.Run()

	go func() {
		var s string
		fmt.Scanln(&s)
		close(done)
		//sig := <-sigs
		//fmt.Println()
		//fmt.Println(sig)
		//close(done)
	}()

	<-done
	s.Shutdown(ctx)
}
