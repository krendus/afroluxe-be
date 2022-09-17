package models

import "mime/multipart"

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Image struct {
	Id        string `json:"_id,omitempty" bson:"_id"`
	Type      string `json:"type" bson:"type"`
	Url       string `json:"url" bson:"url"`
	StylistId string `json:"-" bson:"stylist_id"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}
