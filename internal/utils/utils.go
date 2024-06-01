package utils

import (
	"fidelizou-go/internal/entities/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSession(c *gin.Context) (session models.Session) {
	sessionAny, exists := c.Get("session")

	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "session not found"})
		return
	}

	session = sessionAny.(models.Session)
	return session
}

func WriteError(c *gin.Context, err error) {

	c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})

}
