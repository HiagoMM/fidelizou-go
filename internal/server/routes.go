package server

import (
	"fidelizou-go/internal/db"
	"fidelizou-go/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	OnlyProvider      = []db.UserRole{db.UserRolePROVIDER}
	ProviderAndClient = []db.UserRole{db.UserRoleCLIENT, db.UserRolePROVIDER}
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()
	h := handlers.NewHandlers(s.db)

	r.Use(ErrorMiddleware())

	base := r.Group("/")
	base.GET("/auth/:provider", h.AuthLoginHandler)
	base.GET("/logout/:provider", h.AuthLogoutHandler)
	base.GET("/auth/:provider/callback", h.AuthCallbackHandler)

	user := base.Group("/users", AuthMiddleware(ProviderAndClient))
	user.GET("/me", h.GetUserHandler)
	user.PATCH("/toggle-role", h.ToggleUserRoleHandler)

	return r
}
