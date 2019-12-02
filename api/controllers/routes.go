package controllers

import "github.com/nidhinp/todo/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.GET("/", middlewares.SetJSONMiddleware(s.HomeController))
	s.Router.POST("/login", middlewares.SetJSONMiddleware(s.Login))
}
