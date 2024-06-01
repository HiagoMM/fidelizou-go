package server

import (
	"fidelizou-go/internal/entities/models"
	"fidelizou-go/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	OnlyProvider      = []int{models.RoleClient}
	ProviderAndClient = []int{models.RoleClient, models.RoleProvider}
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	h := handlers.NewHandlers(s.db)

	r.Use(ErrorMiddleware())

	r.GET("/health", h.HealthHandler)

	base := r.Group("/")
	base.POST("/login", h.LoginHandler)

	user := base.Group("/users", AuthMiddleware(ProviderAndClient))
	user.GET("/me", h.GetUserHandler)
	user.PATCH("/toggle-role", h.ToggleUserRoleHandler)

	return r
}
