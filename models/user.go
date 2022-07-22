package models

type User struct {
	FirstName       string  `json:"first_name" bson:"first_name" binding:"required"`
	LastName        string  `json:"last_name" bson:"last_name" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Password        string  `json:"password" binding:"required"`
	MobileNumber    string  `json:"mobile_number" bson:"mobile_number" binding:"required"`
	ResidentCountry string  `json:"resident_country" bson:"resident_country,omitempty" binding:"required"`
	Longitude       float64 `bson:"longitude,omitempty"`
	Latitude        float64 `bson:"latitude,omitempty"`
	Role            string  `binding:"required"`
}
type Review struct {
	Body                 string
	ServiceRating        uint8  `json:"service_rating" bson:"service_rating"`
	RecommendationRating uint8  `json:"recommendation_rating" bson:"recommendation_rating"`
	PricingRating        uint8  `json:"pricing_rating" bson:"pricing_rating"`
	StylistId            string `json:"stylist_id" bson:"stylist_id"`
}
