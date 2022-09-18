package models

// Review stylist review struct
type Review struct {
	Id                   string `json:"_id" bson:"_id,omitempty"`
	Body                 string `json:"body" bson:"body" binding:"required"`
	ServiceRating        uint8  `json:"service_rating" bson:"service_rating" binding:"required"`
	RecommendationRating uint8  `json:"recommendation_rating" bson:"recommendation_rating" binding:"required"`
	PricingRating        uint8  `json:"pricing_rating" bson:"pricing_rating" binding:"required"`
	StylistId            string `json:"stylist_id" bson:"stylist_id" binding:"required"`
	UserId               string `json:"user_id" bson:"user_id" binding:"required"`
	CreatedAt            int64  `json:"created_at" bson:"created_at"`
}

type Location struct {
	Longitude float64 `json:"longitude" bson:"longitude" binding:"required"`
	Latitude  float64 `json:"latitude" bson:"latitude" binding:"required"`
}

type BusinessHour struct {
	Sunday    string `json:"sunday" bson:"sunday" binding:"required"`
	Monday    string `json:"monday" bson:"monday" binding:"required"`
	Tuesday   string `json:"tuesday" bson:"tuesday" binding:"required"`
	Wednesday string `json:"wednesday" bson:"wednesday" binding:"required"`
	Thursday  string `json:"thursday" bson:"thursday" binding:"required"`
	Friday    string `json:"friday" bson:"friday" binding:"required"`
	Saturday  string `json:"saturday" bson:"saturday" binding:"required"`
}

type Service struct {
	Id             string  `json:"_id" bson:"_id"`
	Name           string  `json:"name" bson:"name" binding:"required"`
	Price          float64 `json:"price" bson:"price" binding:"required"`
	CurrencyName   string  `json:"currency_name" bson:"currency_name" binding:"required"`
	CurrencySymbol string  `json:"currency_symbol" bson:"currency_symbol" binding:"required"`
	StylistId      string  `json:"stylist_id" bson:"stylist_id"`
}

type Stylist struct {
	Id           string       `json:"_id" bson:"_id,omitempty"`
	BusinessName string       `json:"business_name" bson:"business_name"  binding:"required"`
	Bio          string       `json:"bio" bson:"bio"  binding:"required"`
	UserId       string       `json:"user_id" bson:"user_id"  binding:"required"`
	Location     Location     `json:"location" bson:"location"  binding:"required"`
	BusinessHour BusinessHour `json:"business_hour" bson:"business_hour"  binding:"required"`
	Services     []Service    `json:"services" bson:"-"  binding:"required"`
	Category     []string     `json:"category" bson:"category"  binding:"required"`
	Images       []Image      `json:"images" bson:"-"`
	BusinessType string       `json:"business_type" bson:"business_type"  binding:"required"`
	HomeService  bool         `json:"home_service" bson:"home_service"  binding:"required"`
	Size         uint         `json:"size" bson:"size"  binding:"required"`
	CreatedAt    int64        `json:"created_at" bson:"created_at"`
	UpdatedAt    int64        `json:"updated_at" bson:"updated_at"`
}
