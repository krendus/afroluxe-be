package utils

import (
	"context"
	"time"

	"github.com/afroluxe/afroluxe-be/config"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

var loadedEnv = config.LoadEnv()

func ImageUploader(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(loadedEnv.CloudinaryCloudName, loadedEnv.CloudinaryApiKey, loadedEnv.CloudinaryApiSecret)

	if err != nil {
		return "", err
	}
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: loadedEnv.CloudinaryUploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
