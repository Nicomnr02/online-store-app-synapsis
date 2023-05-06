package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	GetProductByID(id int) (model.Product, error)
	GetProductsByCategoryID(id int) ([]model.Product, error)
	GetAllProducts() ([]model.Product, error)
	StoreManyProducts(products *[]model.Product) ([]model.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) GetProductByID(id int) (model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) GetProductsByCategoryID(id int) ([]model.Product, error) {
	var products []model.Product
	if err := pr.db.Find(&products, "category_id = ?", id).Error; err != nil {
		return []model.Product{}, err
	}

	return products, nil
}

func (pr *ProductRepository) StoreManyProducts(products *[]model.Product) ([]model.Product, error) {

	var newProducts = products
	if err := pr.db.Save(&products).Error; err != nil {
		return []model.Product{}, err
	}

	return *newProducts, nil

}

func (pr *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	if res := pr.db.Find(&products); res.Error != nil {
		return []model.Product{}, res.Error
	}

	return products, nil

}
