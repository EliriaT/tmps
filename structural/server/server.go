package server

import (
	"fmt"
	"io"
	"net/http"
)

// SimpleServer is a simple http server
type SimpleServer struct {
	Msg string
}

func (s *SimpleServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.Msg)
}

// WithLog is a decorator which adds logging to http handlers
type WithLog struct {
	Handler http.Handler
	Logger  io.Writer
}

func (s *WithLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(s.Logger, "Got request with uri: %s \n", r.RequestURI)
	s.Handler.ServeHTTP(w, r)
}
