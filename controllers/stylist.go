package controllers

import (
	"fmt"
	"net/http"

	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
)

func GetStylist(c *gin.Context) {
	token, _ := c.Cookie("token")
	stylistId := c.Param("id")
	verified, id := utils.VerifyToken(token)
	if verified {
		c.JSON(http.StatusOK, fmt.Sprintf("%v, %v", id, stylistId))
		return
	}
	c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
}
