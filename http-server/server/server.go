package server

import (
	"context"
	"fmt"
	"net/http"
)

// defining a server structure
// port -> which ports does the server listens at
/*
server -> takes a pointer to in-build http server
The reason for using the pointer is -
(i) if we shutdown the server using server.Close() we need to essentially mutate the
	actual running server instead of creating a copy of the server and shutting that
(ii) Working with standard library requires pointer
*/
// Router -> takes the pointer to the Router we created, using pointer because we are mutating it
// middlewares -> Slice of middleware structs which are essentially Handler function wrappers
//				  which will be wrapped on top of Handler
type Server struct {
	port        int
	server      *http.Server
	Router      *Router
	middlewares []Middleware
}

// creates a new server instance and returns the pointer because we need to mutate it
// also we do not want many server in case it is called up multiple times
func NewServer(port int) *Server {
	s := &Server{
		port:        port,
		middlewares: []Middleware{},
	}

	s.Router = NewRouter(s)
	return s
}

// registers all the middlewares 
func (s *Server) Use(middleware ...Middleware) {
	s.middlewares = append(s.middlewares, middleware...)
}

// wraps the Middleware functions on top of Handler
// Handler -> LoggingMiddleware(RecoveryMiddleware(Handler(w, r)))
func (s *Server) ApplyMiddleware(h Handler) Handler {
	if len(s.middlewares) == 0 {
		return h
	}
	return Chain(h, s.middlewares...)
}

// Necessary function to register it as the server in http.Server
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

// Starts the http server in the specified port
func (s *Server) Run() error {
	addr := fmt.Sprintf(":%d", s.port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s,
	}
	fmt.Printf("Server started on port %d\n", s.port)
	return s.server.ListenAndServe()
}


// ShutDown gracefully stops the server with a timeout
// waiting for ongoing requests to be completed before shutting down
// In case the request takes too long to respond the server shutsdown after the context timeout
func (s *Server) ShutDown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}