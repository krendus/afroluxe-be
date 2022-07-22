package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to afroluxe api")
}
