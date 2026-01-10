package server

import "net/http"

// Middleware functions are just the wrapper on top of Handler func(w, r)
type Middleware func(Handler) Handler

// Wraps all the middleware function on top of handler functions, ordering can matter
// here we are wrapping the Handler function in rev order acc to our implementation
// chain(h, LoggingMilddleware, RecoveryMiddleware)
// Logging(Recovery(Handler))
func Chain(h Handler, middleware ...Middleware) Handler {
	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}
	return h
}

// Simply logs the incoming request and response
func LoggingMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		println("Request: ", r.Method, r.URL.Path)
		next(w, r)
		println("Response Sent")
	}
}

// Returns a 500 internal server error if a panic occurs
func RecoveryMiddleware(next Handler) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Panic occured, recover and send error response
				println("Panic recovered: ", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			} 
		}()
		next(w, r)
	}
}