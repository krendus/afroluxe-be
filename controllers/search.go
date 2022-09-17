package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	searchQuery := c.Query("q")
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("search result for %v", searchQuery)})
}
