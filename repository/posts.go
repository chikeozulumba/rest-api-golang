package repository

import "api/models"

type PostsRepository interface {
	Save(models.Post) (models.Post, error)
	FindAll() ([]models.Post, error)
	FindById(uint64) (models.Post, error)
	Update(uint64, models.Post) (int64, error)
	DeletePost(uint64) (int64, error)
}