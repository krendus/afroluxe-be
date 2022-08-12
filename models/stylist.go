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
	Longitude float64 `json:"longitude" bson:"longitude"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
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
	Name         string `json:"name" bson:"name"  binding:"required"`
	Bio          string `json:"bio" bson:"bio"  binding:"required"`
	UserId       string `json:"user_id" bson:"user_id"  binding:"required"`
	Location     `json:"location" bson:"location"  binding:"required"`
	BusinessHour `json:"business_hour" bson:"business_hour"  binding:"required"`
	Services     []string `json:"services" bson:"services"  binding:"required"`
	Category     []string `json:"category" bson:"category"  binding:"required"`
	Type         string   `json:"type" bson:"type"  binding:"required"`
	HomeService  bool     `json:"home_service" bson:"home_service"  binding:"required"`
	Size         uint     `json:"size" bson:"size"  binding:"required"`
	CreatedAt    int64    `json:"created_at" bson:"created_at"`
	UpdatedAt    int64    `json:"updated_at" bson:"updated_at"`
}
