package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func RepositoryProduct(db *gorm.DB) *productRepo {
	return &productRepo{db}
}


func (r *repo) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Raw("SELECT * FROM products").Scan(&products).Error

	return products, err
}