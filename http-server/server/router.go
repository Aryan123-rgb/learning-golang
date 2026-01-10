package server

import "net/http"

// the necessary skeleton for the handler function to process http request, response
type Handler func(w http.ResponseWriter, r *http.Request)

type Router struct {
	routes map[string]map[string]Handler // method (POST, GET) -> path -> handler
	server *Server // the pointer to the server 
}

// returns the pointer to the new Router we created
func NewRouter(s *Server) *Router {
	return &Router{
		routes: make(map[string]map[string]Handler),
		server: s,
	}
}

// registers all the route
func (r *Router) addRoute(method, path string, handler Handler) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]Handler)
	}
	r.routes[method][path] = handler
}

// registers all the GET requests using addRoute function
func (r *Router) GET(path string, handler Handler) {
	r.addRoute(http.MethodGet, path, handler)
}

// registers all the POST requests using addRoute function
func (r *Router) POST(path string, handler Handler) {
	r.addRoute(http.MethodPost, path, handler)
}

// function that the golang router invokes internally after recieveing any request
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// check if the given method and route exists
	if methodRoutes, ok := r.routes[req.Method]; ok {
		if handler, ok := methodRoutes[req.URL.Path]; ok {

			// apply the middlewares
			if r.server != nil {
				handler = r.server.ApplyMiddleware(handler)
			}

			// apply the actual handler (controller function)
			handler(w, req)
			return
		}
	}
	http.Error(w, "404 page not found", http.StatusNotFound)
}
