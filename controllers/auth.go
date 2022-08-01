package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection *mongo.Collection = db.CollectionInstance("users")

type LoginCredentials struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type RegResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

type VerifyRequest struct {
	Otp   string `json:"otp" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// HandleSignUp handles the signup logic
func HandleSignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required fields are missing"})
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

		// save the user struct on redis
		err = db.SetRedisValue(user.Email, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
			return
		}

		// generates random otp with a length of 6
		otp := utils.GenerateRandomOtp(6)

		// saves the otp to redis
		err = db.SetRedisValue(fmt.Sprintf("%v-otp", user.Email), otp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
			return
		}

		// sends mail to user
		err := utils.SendEmail(utils.Mail{
			From:     "samuellawal1979@gmail.com",
			To:       []string{user.Email},
			Subject:  "OTP to verify your email",
			Data:     struct{ Otp string }{otp},
			Filename: "./template/verify.html",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		}
		c.JSON(http.StatusOK, RegResponse{"OTP is sent to email for verification", user.Email})
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

	// validates login
	err := UserCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
	if err != mongo.ErrNoDocuments {
		if valid := utils.CheckPasswordHash(result.Password, user.Password); valid {
			expTime := time.Now().Add(time.Minute * 60 * 3)
			token, err := utils.CreateNewToken(result.Id, expTime)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				return
			}
			c.SetCookie("token", token, 60*60*3, "/", "*", false, true)
			c.JSON(http.StatusOK, result.Res())
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid login details"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid login details"})
	}

}

func VerifyEmail(c *gin.Context) {
	var body VerifyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "required fields are missing"})
		return
	}
	var redisOtp string

	// fetch the otp from redis
	err := db.GetRedisValue(fmt.Sprintf("%v-otp", body.Email), &redisOtp)
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Registration session expired"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		return
	}

	// comapares the sent otp and the redis
	if redisOtp != body.Otp {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OTP"})
		return
	}
	var user models.User
	err = db.GetRedisValue(body.Email, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		return
	}
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		return
	}
	err = db.DelRedisValue(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		return
	}
	err = db.DelRedisValue(fmt.Sprintf("%v-otp", user.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal sever error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Registration succesful"})
}
