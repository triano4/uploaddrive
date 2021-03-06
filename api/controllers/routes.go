package controllers

import (
	"github.com/triano4/uploaddrive/api/middlewares"
	"github.com/triano4/uploaddrive/code"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST", "OPTIONS")

	// s.Router.HandleFunc("/login", s.Login).Methods("POST")

	s.Router.HandleFunc("/test", s.EndpointTest)

	// s.Router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Test") })
	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Other routes
	s.Router.HandleFunc("/file", middlewares.SetMiddlewareJSON(code.Client))

}
