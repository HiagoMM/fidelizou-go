package utils

import (
	"fidelizou-go/internal/entities/err"
	"fidelizou-go/internal/entities/models"

	"github.com/gin-gonic/gin"
)

var keySession = "session"

func GetSession(c *gin.Context) (session models.Session) {
	sessionAny, exists := c.Get(keySession)

	if !exists {
		c.Error(err.NewUnauthorized("Sessão não encontrada"))
		c.Abort()
		return
	}

	session = sessionAny.(models.Session)
	return session
}

func SetSession(c *gin.Context, session models.Session) {
	c.Set(keySession, session)
}
