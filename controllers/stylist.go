package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var stylistCollection = db.CollectionInstance("stylists")

func GetStylist(c *gin.Context) {
	token, _ := c.Cookie("token")
	stylistId := c.Param("id")
	verified, _ := utils.VerifyToken(token)
	if !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	var result models.Stylist

	mongoStylistId, err := primitive.ObjectIDFromHex(stylistId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	err = stylistCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: mongoStylistId}}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "stylist not found", "data": result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func CreateStylist(c *gin.Context) {
	var stylist models.Stylist
	token, _ := c.Cookie("token")
	verified, id := utils.VerifyToken(token)
	if !verified {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	if err := c.ShouldBindJSON(&stylist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "required fields are missing"})
		return
	}
	if id != stylist.UserId {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	err := stylistCollection.FindOne(context.TODO(), bson.D{{Key: "user_id", Value: stylist.UserId}}).Err()

	if err == mongo.ErrNoDocuments {
		stylist.CreatedAt = time.Now().Unix()
		stylist.UpdatedAt = time.Now().Unix()
		_, err = stylistCollection.InsertOne(context.TODO(), stylist)
		if err != nil {
			utils.ErrorLogger(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal sever error"})
		}

		c.JSON(http.StatusCreated, gin.H{"message": "successfully created stylist"})
		return
	}
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	utils.ErrorLogger(err)
	c.JSON(http.StatusBadRequest, gin.H{"message": "Stylist already exists"})
}
