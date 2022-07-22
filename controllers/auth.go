package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = db.CollectionInstance("user")

func HandleSignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required fields are missing"})
	} else if user.Role == "stylist" && user.Latitude == 0 && user.Longitude == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stylist location must be provided"})
	} else {
		var result bson.M
		err := UserCollection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}}).Decode(&result)
		if err == mongo.ErrNoDocuments {
			user.Password, _ = utils.HashPassword(user.Password)
			res, err := UserCollection.InsertOne(context.TODO(), user)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(http.StatusCreated, res)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already taken"})
		}
	}
}
func HandleSignIn(c *gin.Context) {
	c.JSON(http.StatusOK, "working")
}
