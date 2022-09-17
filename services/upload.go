package services

import (
	"github.com/afroluxe/afroluxe-be/models"
	"github.com/afroluxe/afroluxe-be/utils"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file models.File) (string, error)
}

type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (m *media) FileUpload(file models.File) (string, error) {
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}
	uploadUrl, err := utils.ImageUploader(file.File)
	if err != nil {
		return "", err
	}

	return uploadUrl, nil
}
