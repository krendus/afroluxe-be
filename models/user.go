package models

type User struct {
	Id              string  `json:"_id" bson:"_id,omitempty"`
	FirstName       string  `json:"first_name" bson:"first_name" binding:"required"`
	LastName        string  `json:"last_name" bson:"last_name" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Password        string  `json:"password" binding:"required"`
	MobileNumber    string  `json:"mobile_number" bson:"mobile_number" binding:"required"`
	ResidentCountry string  `json:"resident_country" bson:"resident_country,omitempty" binding:"required"`
	Longitude       float64 `json:"longitude" bson:"longitude"`
	Latitude        float64 `json:"latitude" bson:"latitude"`
	Role            string  `json:"role" binding:"required"`
}
type Review struct {
	Id                   string `json:"_id" bson:"_id,omitempty"`
	Body                 string
	ServiceRating        uint8  `json:"service_rating" bson:"service_rating"`
	RecommendationRating uint8  `json:"recommendation_rating" bson:"recommendation_rating"`
	PricingRating        uint8  `json:"pricing_rating" bson:"pricing_rating"`
	StylistId            string `json:"stylist_id" bson:"stylist_id"`
}
