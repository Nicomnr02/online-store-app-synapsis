package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
)

type CategoriesAPIInterface interface{
	StoreManyCategories(w http.ResponseWriter, r *http.Request)
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	GetAllCategoriesWithProducts(w http.ResponseWriter, r *http.Request)
}

type CategoriesAPI struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryAPI(categoryService services.CategoryServiceInterface) *CategoriesAPI {
	return &CategoriesAPI{categoryService}
}

func (ca *CategoriesAPI) StoreManyCategories(w http.ResponseWriter, r *http.Request) {
	var categories []model.Category
	err := json.NewDecoder(r.Body).Decode(&categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	if newCategories, err := ca.categoryService.StoreManyCategories(categories); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success store many categories", Data: newCategories})
	}

}

func (ca *CategoriesAPI) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	if categories, err := ca.categoryService.GetAllCategories(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get all categories", Data: categories})
	}
}

func (ca *CategoriesAPI) GetAllCategoriesWithProducts(w http.ResponseWriter, r *http.Request) {
	if categoriesWithProducts, err := ca.categoryService.GetAllCategoriesWithProducts(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get categories with products", Data: categoriesWithProducts})
	}

}
