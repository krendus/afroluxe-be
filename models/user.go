package models

import "time"

// User user model
type User struct {
	Id           string    `json:"_id" bson:"_id,omitempty"`
	FirstName    string    `json:"first_name" bson:"first_name" binding:"required"`
	LastName     string    `json:"last_name" bson:"last_name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	Password     string    `json:"password" binding:"required"`
	MobileNumber string    `json:"mobile_number" bson:"mobile_number" binding:"required"`
	Role         string    `json:"role" binding:"required"`
	Joined       time.Time `json:"joined"`
}
type UserRes struct {
	Id           string    `json:"_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	MobileNumber string    `json:"mobile_number"`
	Role         string    `json:"role"`
	Joined       time.Time `json:"joined"`
}

func (u User) Res() UserRes {
	return UserRes{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		MobileNumber: u.MobileNumber,
		Role:         u.Role,
		Joined:       u.Joined,
	}
}
