package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

//Handler struct holds required services for handler to function
type Handler struct {}

// Config will hold services that will be injected to this handler layer on handler initialization
type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config) {
	h := &Handler{}

	g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))
	g.GET("/signup", h.Signup)
}

// Signup Sign up handler
func (h *Handler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
		"method": "Sign up",
	})
}