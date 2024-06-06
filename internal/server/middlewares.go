package server

import (
	"errors"
	"fidelizou-go/internal/db"
	"fidelizou-go/internal/entities/err"
	"fidelizou-go/internal/utils"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var httpError *err.HttpError
		for _, e := range c.Errors {
			switch {
			case errors.As(e.Err, &httpError):
				c.JSON(httpError.Status, gin.H{"message": httpError.Error()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"message": e.Error()})
			}
		}
		c.Next()
	}
}

func AuthMiddleware(canAccess []db.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := utils.GetSession(c)
		if !slices.Contains(canAccess, db.UserRole(session.Role)) {
			c.Error(err.NewForbidden("Usuário não autorizado"))
			c.Abort()
			return
		}
		c.Next()
	}
}
