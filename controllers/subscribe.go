package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/dtos"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var subscriptionCollection = db.CollectionInstance("subscription")

// Subcribe ... Adds email to subscription list
// @Summary Adds email to subscription list
// @Description Adds email to subscription list
// @Tags Subscribe
// @Accept json
// @Param subscribe body models.VerifyRequest true "Email"
// @Success 200 {object} dtos.Response
// @Failure 400,500 {object} dtos.Response
// @Router /subscribe [post]
func Subscribe(c *gin.Context) {
	var subscription models.Subscribe

	if err := c.ShouldBindJSON(&subscription); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "required fields are missing",
			Data:       nil,
		})
		return
	}

	err := subscriptionCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: subscription.Email}}).Err()

	if err == mongo.ErrNoDocuments {
		subscription.CreatedAt = time.Now().Unix()
		_, err = subscriptionCollection.InsertOne(context.TODO(), subscription)
		if err != nil {
			utils.ErrorLogger(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "internal server error",
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusCreated, dtos.Response{
			StatusCode: http.StatusCreated,
			Message:    "email added to subscripion list",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusBadRequest, dtos.Response{
		StatusCode: http.StatusBadRequest,
		Message:    "email already added",
		Data:       nil,
	})
}
