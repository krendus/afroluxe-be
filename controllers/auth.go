package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = db.CollectionInstance("users")

type LoginCredentials struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// HandleSignUp handles the signup logic
func HandleSignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required fields are missing"})
		return
	}
	if user.Role == "stylist" && user.Latitude == 0 && user.Longitude == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stylist location must be provided"})
		return
	}
	if user.Role != "stylist" && user.Role != "customer" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user role"})
		return
	}
	var result bson.M
	err := UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		user.Password, _ = utils.HashPassword(user.Password)
		user.Joined = time.Now()
		res, err := UserCollection.InsertOne(context.TODO(), user)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusCreated, res)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken"})
	}
}

// HandleSignIn Handle the signin logic
func HandleSignIn(c *gin.Context) {
	var user LoginCredentials
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required fields are missing"})
		return
	}
	var result models.User
	err := UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
	if err != mongo.ErrNoDocuments {
		if valid := utils.CheckPasswordHash(result.Password, user.Password); valid {
			c.JSON(http.StatusOK, result.Res())
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login details"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login details"})
	}

}
