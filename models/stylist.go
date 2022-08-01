package models

import "time"

// Review stylist review struct
type Review struct {
	Id                   string    `json:"_id" bson:"_id,omitempty"`
	Body                 string    `json:"body" bson:"body"`
	ServiceRating        uint8     `json:"service_rating" bson:"service_rating"`
	RecommendationRating uint8     `json:"recommendation_rating" bson:"recommendation_rating"`
	PricingRating        uint8     `json:"pricing_rating" bson:"pricing_rating"`
	StylistId            string    `json:"stylist_id" bson:"stylist_id"`
	CustomerId           string    `json:"customer_id" bson:"customer_id"`
	CreatedAt            time.Time `json:"created_at" bson:"created_at"`
}

// Service stylist service struct
type Service struct {
	Id        string    `json:"_id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Price     uint      `json:"price" bson:"price"`
	StylistId string    `json:"stylist_id" bson:"stylist_id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type Location struct {
	Longitude string `json:"longitude" bson:"longitude"`
	Latitude  string `json:"latitude" bson:"latitude"`
}

type BusinessHour struct {
	Sunday    string `json:"sunday" bson:"sunday"`
	Monday    string `json:"monday" bson:"monday"`
	Tuesday   string `json:"tuesday" bson:"tuesday"`
	Wednesday string `json:"wednesday" bson:"wednesday"`
	Thursday  string `json:"thursday" bson:"thursday"`
	Friday    string `json:"friday" bson:"friday"`
	Saturday  string `json:"saturday" bson:"saturday"`
}

type Stylist struct {
	Id           string `json:"_id" bson:"_id,omitempty"`
	Name         string `json:"name" bson:"name"`
	Bio          string `json:"bio" bson:"bio"`
	UserId       string `json:"user_id" bson:"user_id"`
	Location     `json:"location" bson:"location"`
	BusinessHour `json:"business_hour" bson:"business_hour"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
}
