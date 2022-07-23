package models

import "time"

type User struct {
	Id              string    `json:"_id" bson:"_id,omitempty"`
	FirstName       string    `json:"first_name" bson:"first_name" binding:"required"`
	LastName        string    `json:"last_name" bson:"last_name" binding:"required"`
	Email           string    `json:"email" binding:"required"`
	Password        string    `json:"password" binding:"required"`
	MobileNumber    string    `json:"mobile_number" bson:"mobile_number" binding:"required"`
	ResidentCountry string    `json:"resident_country" bson:"resident_country,omitempty" binding:"required"`
	Longitude       float64   `json:"longitude" bson:"longitude"`
	Latitude        float64   `json:"latitude" bson:"latitude"`
	Role            string    `json:"role" binding:"required"`
	Joined          time.Time `json:"joined"`
}
type UserRes struct {
	Id              string    `json:"_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	MobileNumber    string    `json:"mobile_number"`
	ResidentCountry string    `json:"resident_country"`
	Longitude       float64   `json:"longitude"`
	Latitude        float64   `json:"latitude"`
	Role            string    `json:"role"`
	Joined          time.Time `json:"joined"`
}

func (u User) Res() UserRes {
	return UserRes{
		Id:              u.Id,
		FirstName:       u.FirstName,
		LastName:        u.LastName,
		Email:           u.Email,
		MobileNumber:    u.MobileNumber,
		ResidentCountry: u.ResidentCountry,
		Longitude:       u.Longitude,
		Latitude:        u.Latitude,
		Role:            u.Role,
		Joined:          u.Joined,
	}
}

type Review struct {
	Id                   string `json:"_id" bson:"_id,omitempty"`
	Body                 string
	ServiceRating        uint8  `json:"service_rating" bson:"service_rating"`
	RecommendationRating uint8  `json:"recommendation_rating" bson:"recommendation_rating"`
	PricingRating        uint8  `json:"pricing_rating" bson:"pricing_rating"`
	StylistId            string `json:"stylist_id" bson:"stylist_id"`
}
