package repositories

import (
	"backend/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindProducts() ([]models.Product, error)
	GetProducts(ID int) (models.Product, error)
	CreateProduct(product models.Product) (models.Product, error)
	UpdateProduct(product models.Product, ID int) (models.Product, error)
	DeleteProduct(product models.Product, ID int) (models.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func RepositoryProduct(db *gorm.DB) *productRepo {
	return &productRepo{db}
}


func (r *productRepo) FindProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Find(&products).Error // Using Find method

	return products, err
}

func (r *productRepo) GetProducts(ID int) (models.Product, error) {
	var products models.Product
	err := r.db.First(&products, ID).Error // Using First methoderr := r.db.First(&user, ID).Error // Using First method

	return products, err
}

func (r *productRepo) CreateProduct(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error // Using Create method ORM

	return product, err
}

func (r *productRepo) UpdateProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Save(&product).Error // Using Save method ORM

	return product, err
}

func (r *productRepo) DeleteProduct(product models.Product, ID int) (models.Product, error) {
	err := r.db.Delete(&product).Error // Using Delete method ORM

	return product, err
}