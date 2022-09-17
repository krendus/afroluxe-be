package models

type Subscribe struct {
	Id        string `json:"_id" bson:"_id,omitempty"`
	Email     string `json:"email" bson:"email" binding:"required"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}
