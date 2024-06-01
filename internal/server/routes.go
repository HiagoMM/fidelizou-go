package server

import (
	"fidelizou-go/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	h := handlers.NewHandlers(s.db)

	r.GET("/health", h.HealthHandler)

	user := r.Group("/user")
	user.GET("/me", h.GetUserHandler)

	return r
}
