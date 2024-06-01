package handlers

import (
	"fidelizou-go/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) GetUserHandler(c *gin.Context) {
	session := utils.GetSession(c)

	user, err := h.s.GetUser(session.Email)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handlers) ToggleUserRoleHandler(c *gin.Context) {
	session := utils.GetSession(c)

	if err := h.s.ToggleUserRole(&session); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}
