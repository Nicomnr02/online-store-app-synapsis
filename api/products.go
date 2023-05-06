package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
	"strconv"
)

type ProductAPIInterface interface {
	StoreManyProducts(w http.ResponseWriter, r *http.Request)
	GetProductsByCategoryID(w http.ResponseWriter, r *http.Request)
	GetProductsByID(w http.ResponseWriter, r *http.Request)
	GetAllProducts(w http.ResponseWriter, r *http.Request)
}

type ProductAPI struct {
	productService services.ProductServiceInterface
}

func NewProductAPI(productService services.ProductServiceInterface) *ProductAPI {
	return &ProductAPI{productService}
}

func (pa *ProductAPI) StoreManyProducts(w http.ResponseWriter, r *http.Request) {
	var products []model.Product
	err := json.NewDecoder(r.Body).Decode(&products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	if newProducts, err := pa.productService.StoreManyProducts(products); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success store many categories", Data: newProducts})
	}
}

func (pa *ProductAPI) GetProductsByCategoryID(w http.ResponseWriter, r *http.Request) {
	catID, _ := strconv.Atoi(r.URL.Query().Get("category_id"))

	if products, err := pa.productService.GetProductsByCategoryID(catID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get products by category id", Data: products})
	}
}

func (pa *ProductAPI) GetProductsByID(w http.ResponseWriter, r *http.Request) {
	prodID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))

	if product, err := pa.productService.GetProductByID(prodID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get product by id", Data: product})
		return
	}
}

func (pa *ProductAPI) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	if products, err := pa.productService.GetAllProducts(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get all products", Data: products})
	}
}
