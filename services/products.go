package services

import (
	"online_app_store/model"
	"online_app_store/repositories"
)

type ProductServiceInterface interface{
	StoreManyProducts(products []model.Product) ([]model.Product, error)
	GetProductsByCategoryID(id int) ([]model.Product, error)
	GetProductByID(id int) (model.Product, error)
	GetAllProducts() ([]model.Product, error)
}

type ProductService struct {
	productRepo repositories.ProductRepositoryInterface
}

func NewProductService(productRepo repositories.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{productRepo}
}

func (ps *ProductService) StoreManyProducts(products []model.Product) ([]model.Product, error) {
	if newProducts, err := ps.productRepo.StoreManyProducts(&products); err != nil {
		return []model.Product{}, err
	} else {
		return newProducts, nil
	}
}

func (ps *ProductService) GetProductsByCategoryID(id int) ([]model.Product, error) {
	if products, err := ps.productRepo.GetProductsByCategoryID(id); err != nil {
		return []model.Product{}, err
	} else {
		return products, nil
	}
}

func (ps *ProductService) GetProductByID(id int) (model.Product, error) {
	if product, err := ps.productRepo.GetProductByID(id); err != nil {
		return model.Product{}, err
	} else {
		return product, nil
	}
}

func (ps *ProductService) GetAllProducts() ([]model.Product, error) {
	if products, err := ps.productRepo.GetAllProducts(); err != nil {
		return []model.Product{}, err
	} else {
		return products, nil
	}
}
