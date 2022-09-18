package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/afroluxe/afroluxe-be/config"
	"github.com/afroluxe/afroluxe-be/db"
	"github.com/afroluxe/afroluxe-be/dtos"
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/services"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection = db.CollectionInstance("users")
	loadedEnv      = config.LoadEnv()
)

// HandleSignUp ... Create User
// @Summary Create new user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body models.User true "User Data"
// @Success 200 {object} dtos.Response
// @Failure 400,500 {object} dtos.Response
// @Router /auth/signup [post]
func HandleSignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Required fields are missing",
			Data:       nil,
		})
		return
	}
	if user.Role != "stylist" && user.Role != "customer" {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user role",
			Data:       nil,
		})
		return
	}

	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Err()
	if err != nil && err != mongo.ErrNoDocuments {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	if err == mongo.ErrNoDocuments {
		user.Password, _ = utils.HashPassword(user.Password)
		user.Joined = time.Now().Unix()

		// save the user struct on redis.
		err = db.SetRedisValue(user.Email, user)
		if err != nil {
			utils.ErrorLogger(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
				Data:       nil,
			})
			return
		}

		// generates random otp with a length of 6.
		otp := utils.GenerateRandomOtp(6)

		// saves the otp to redis.
		err = db.SetRedisValue(fmt.Sprintf("%v-otp", user.Email), otp)
		if err != nil {
			utils.ErrorLogger(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
				Data:       nil,
			})
			return
		}

		// sends mail to user.
		err := services.SendEmail(services.Mail{
			From:     "samuellawal1979@gmail.com",
			To:       []string{user.Email},
			Subject:  "OTP to verify your email",
			Data:     struct{ Otp string }{otp},
			Filename: "./template/verify.html",
		})
		if err != nil {
			utils.ErrorLogger(err)
			c.JSON(http.StatusInternalServerError, dtos.Response{
				StatusCode: http.StatusInternalServerError,
				Message:    "Internal server error",
				Data:       nil,
			})
			return
		}
		c.JSON(http.StatusOK, dtos.Response{
			StatusCode: http.StatusOK,
			Message:    "OTP is sent to email for verification",
			Data:       gin.H{"email": user.Email},
		})
	} else {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid user role",
			Data:       nil,
		})
	}
}

// HandleSignIn ... SignIn User
// @Summary Sign in user based on paramters
// @Description Create new user
// @Tags Users
// @Accept json
// @Param user body models.LoginCredentials true "User Data"
// @Success 200 {object} dtos.Response
// @Failure 400,500 {object} dtos.Response
// @Router /auth/signin [post]
func HandleSignIn(c *gin.Context) {
	var user models.LoginCredentials
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Required fields are missing",
			Data:       nil,
		})
		return
	}
	var result models.User

	// validates login
	err := userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: user.Email}}).Decode(&result)
	if err == nil {
		if valid := utils.CheckPasswordHash(result.Password, user.Password); valid {
			token, err := utils.CreateNewToken(result.Id)
			if err != nil {
				utils.ErrorLogger(err)
				c.JSON(http.StatusInternalServerError, dtos.Response{
					StatusCode: http.StatusInternalServerError,
					Message:    "Internal server error",
					Data:       nil,
				})
				return
			}
			duration, err := strconv.Atoi(loadedEnv.JwtDuration)
			if err != nil {
				utils.ErrorLogger(err)
			}
			c.SetCookie("token", token, duration, "/", "", false, true)
			c.JSON(http.StatusOK, dtos.Response{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       result.Res(),
			})
		} else {
			c.JSON(http.StatusUnauthorized, dtos.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid login details",
				Data:       nil,
			})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, dtos.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid login details",
			Data:       nil,
		})
	}

}

// VerifyEmail ... Verifies user email
// @Summary Verifies user based on paramters
// @Description Verifies user email
// @Tags Users
// @Accept json
// @Param user body models.VerifyRequest true "OTP"
// @Success 200 {object} dtos.Response
// @Failure 400,500 {object} dtos.Response
// @Router /auth/verify [post]
func VerifyEmail(c *gin.Context) {
	var body models.VerifyRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Required fields are missing",
			Data:       nil,
		})
		return
	}
	var redisOtp string

	// fetch the otp from redis
	err := db.GetRedisValue(fmt.Sprintf("%v-otp", body.Email), &redisOtp)
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, dtos.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Registration session expired",
			Data:       nil,
		})
		return
	}
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}

	// comapares the sent otp and the redis
	if redisOtp != body.Otp {
		c.JSON(http.StatusUnauthorized, dtos.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid OTP",
			Data:       nil,
		})
		return
	}
	var user models.User
	err = db.GetRedisValue(body.Email, &user)
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	userRes, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	err = db.DelRedisValue(user.Email)
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	err = db.DelRedisValue(fmt.Sprintf("%v-otp", user.Email))
	if err != nil {
		utils.ErrorLogger(err)
		c.JSON(http.StatusInternalServerError, dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
		})
		return
	}
	c.JSON(http.StatusCreated, dtos.Response{
		StatusCode: http.StatusCreated,
		Message:    "Registration succesful",
		Data:       gin.H{"user_id": userRes.InsertedID},
	})
}
