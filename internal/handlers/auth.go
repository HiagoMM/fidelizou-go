package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func (h *Handlers) AuthLoginHandler(c *gin.Context) {
	fixProviderParam(c)
	log.Print("Entrou")
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		gothic.StoreInSession("session", user.Email, c.Request, c.Writer)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}

func (h *Handlers) AuthLogoutHandler(c *gin.Context) {
	fixProviderParam(c)
	gothic.Logout(c.Writer, c.Request)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h *Handlers) AuthCallbackHandler(c *gin.Context) {
	fixProviderParam(c)
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		return
	}

	gothic.StoreInSession("session", user.Email, c.Request, c.Writer)
	log.Println(user)

	c.Redirect(http.StatusFound, "/")
}

func (h *Handlers) GetSessionHandler(c *gin.Context) {
	session, err := gothic.GetFromSession("session", c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, session)
}

func fixProviderParam(c *gin.Context) {
	q := c.Request.URL.Query()
	q.Add("provider", c.Param("provider"))
	c.Request.URL.RawQuery = q.Encode()
}
