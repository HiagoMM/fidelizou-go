package handlers

import (
	"fidelizou-go/internal/database"
	"fidelizou-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	s services.Service
}

func NewHandlers(repo database.Repository) *Handlers {
	return &Handlers{
		s: services.NewService(repo),
	}
}

func (h *Handlers) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, h.s.Health())
}
