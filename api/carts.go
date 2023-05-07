package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
	"strconv"
)

type CartAPIInterface interface {
	StoreCart(w http.ResponseWriter, r *http.Request)
	GetAllCartsByUserID(w http.ResponseWriter, r *http.Request)
	GetCartByID(w http.ResponseWriter, r *http.Request)
	UpdateCart(w http.ResponseWriter, r *http.Request)
	DeleteCart(w http.ResponseWriter, r *http.Request)
}

type CartAPI struct {
	cartService services.CartServicesInterface
}

func NewCartAPI(cartService services.CartServicesInterface) *CartAPI {
	return &CartAPI{cartService}
}

func (ca *CartAPI) StoreCart(w http.ResponseWriter, r *http.Request) {
	var cart model.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := r.Context().Value("user_id").(float64)
	cart.UserID = int(userID)

	if newCart, err := ca.cartService.StoreCart(cart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success add a cart", Data: newCart})
	}
}

func (ca *CartAPI) GetAllCartsByUserID(w http.ResponseWriter, r *http.Request) {

	userID := int(r.Context().Value("user_id").(float64))

	if carts, err := ca.cartService.GetAllCartsByUserID(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success show all carts by user ID", Data: carts})
	}
}

func (ca *CartAPI) GetCartByID(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value("user_id").(int)
	cartID, _ := strconv.Atoi(r.URL.Query().Get("cart_id"))

	if cart, err := ca.cartService.GetCartByID(cartID, userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success show all carts by user ID", Data: cart})
	}
}

func (ca *CartAPI) UpdateCart(w http.ResponseWriter, r *http.Request) {

	var cart model.Cart
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := r.Context().Value("user_id").(float64)
	cart.UserID = int(userID)

	if cart, err := ca.cartService.UpdateCart(cart); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success update cart", Data: cart})
	}
}

func (ca *CartAPI) DeleteCart(w http.ResponseWriter, r *http.Request) {

	userID := int(r.Context().Value("user_id").(float64))
	cartID, _ := strconv.Atoi(r.URL.Query().Get("cart_id"))

	if err := ca.cartService.RemoveCart(cartID, userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success delete cart"})
	}
}
