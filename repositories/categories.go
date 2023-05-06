package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	GetAllCategories() ([]model.Category, error)
	CreateManyCategories(categories *[]model.Category) ([]model.Category, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (cr *CategoryRepository) GetAllCategories() ([]model.Category, error) {
	var categories []model.Category
	if res := cr.db.Find(&categories); res.Error != nil {
		return []model.Category{}, res.Error
	}

	return categories, nil

}

func (cr *CategoryRepository) CreateManyCategories(categories *[]model.Category) ([]model.Category, error) {
	var newCategories = categories

	if err := cr.db.Save(&newCategories).Error; err != nil {
		return []model.Category{}, err
	}
	return *newCategories, nil
}
