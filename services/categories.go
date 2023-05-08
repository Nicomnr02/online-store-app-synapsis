package services

import (
	"online_app_store/model"
	"online_app_store/repositories"
)

type CategoryServiceInterface interface {
	StoreManyCategories(categories []model.Category) ([]model.Category, error)
	GetAllCategories() ([]model.Category, error)
	GetAllCategoriesWithProducts() ([]model.CategoryData, error)
}

type CategoryService struct {
	categoryRepo repositories.CategoryRepositoryInterface
	productRepo  repositories.ProductRepositoryInterface
}

func NewCategoryService(categoryRepo repositories.CategoryRepositoryInterface, productRepo repositories.ProductRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{categoryRepo, productRepo}
}

func (cs *CategoryService) StoreManyCategories(categories []model.Category) ([]model.Category, error) {
	if categories, err := cs.categoryRepo.CreateManyCategories(&categories); err != nil {
		return []model.Category{}, err
	} else {

		return categories, nil
	}
}

func (cs *CategoryService) GetAllCategories() ([]model.Category, error) {
	if categories, err := cs.categoryRepo.GetAllCategories(); err != nil {
		return []model.Category{}, err
	} else {
		return categories, nil
	}
}

func (cs *CategoryService) GetAllCategoriesWithProducts() ([]model.CategoryData, error) {
	var products []model.Product
	if existedProducts, err := cs.productRepo.GetAllProducts(); err != nil {
		return []model.CategoryData{}, err
	} else {
		products = existedProducts
	}

	var categories []model.Category
	if existedCategory, err := cs.categoryRepo.GetAllCategories(); err != nil {
		return []model.CategoryData{}, err
	} else {
		categories = existedCategory
	}

	var categoryData = CategoriesWithProducts(categories, products)
	return categoryData, nil

}

func CategoriesWithProducts(categories []model.Category, products []model.Product) []model.CategoryData {
	var categoriesData []model.CategoryData

	for _, cats := range categories {
		var productsData []model.Product

		for _, prods := range products {
			if prods.CategoryID == cats.ID {
				productsData = append(productsData, prods)
			}
		}
		categoriesData = append(categoriesData, model.CategoryData{ID: cats.ID, Type: cats.Type, Products: productsData})
	}

	return categoriesData
}
