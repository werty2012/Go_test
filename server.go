package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run() {
	//http.HandleFunc("/get", handleGet)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerRoot)
	mux.HandleFunc("/get", handleGet)

	// Create an HTTP server that listens on port 8000
	http.ListenAndServe(":8000", mux)
}
func handleGet(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, getPostgresData())
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
func (s *Server) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down server....")
	return s.httpServer.Shutdown(ctx)
}

func handlerRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// This prints to STDOUT to show that processing has started
	fmt.Fprint(os.Stdout, "processing request\n")
	// We use `select` to execute a piece of code depending on which
	// channel receives a message first

	select {
	case <-time.After(2 * time.Second):
		// If we receive a message after 2 seconds
		// that means the request has been processed
		// We then write this as the response
		w.Write([]byte("request processed"))
	case <-ctx.Done():
		// If the request gets cancelled, log it
		// to STDERR
		fmt.Fprint(os.Stderr, "request cancelled\n")
	}
}
