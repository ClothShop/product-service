package services

import (
	"github.com/ClothShop/product-service/internal/repositories"
)

type ImageService struct {
	Repo repositories.ImageRepository
}

func NewImageService(repo repositories.ImageRepository) *ImageService {
	return &ImageService{Repo: repo}
}
