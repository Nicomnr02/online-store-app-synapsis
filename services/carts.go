package services

import (
	"online_app_store/model"
	"online_app_store/repositories"
	"time"
)

type CartServicesInterface interface {
	StoreCart(cart model.Cart) (model.Cart, error)
	GetAllCartsByUserID(userID int) (model.CartData, error)
	GetCartByID(id, userID int) (model.Cart, error)
	UpdateCart(cart model.Cart) (model.Cart, error)
	RemoveCart(id, user int) error
}

type CartService struct {
	cartRepository    repositories.CartRepositoryInterface
	productRepository repositories.ProductRepositoryInterface
}

func NewCartService(cartRepository repositories.CartRepositoryInterface, productRepository repositories.ProductRepositoryInterface) CartServicesInterface {
	return &CartService{cartRepository, productRepository}
}

func (cs *CartService) StoreCart(cart model.Cart) (model.Cart, error) { // in live product page

	userID := cart.UserID
	allExistedProductsInCart, err := cs.cartRepository.GetAllCartsByUserID(userID)
	if err != nil {
		return model.Cart{}, err
	}

	if cart.Quantity == 0 {
		cart.Quantity = 1
	}

	productID := cart.ProductID
	product, err := cs.productRepository.GetProductByID(productID)
	if err != nil {
		return model.Cart{}, err
	}

	//cart is filled
	for _, cartData := range allExistedProductsInCart {
		if cartData.UserID == cart.UserID && cartData.ProductID == cart.ProductID {
			cart.ID = cartData.ID
		}
	}
	cart.Price = cart.Quantity * product.Price

	newCart, err := cs.cartRepository.CreateCart(cart)
	if err != nil {
		return model.Cart{}, err
	}

	return newCart, nil
}

func (cs *CartService) GetAllCartsByUserID(userID int) (model.CartData, error) {

	if carts, err := cs.cartRepository.GetAllCartsByUserID(userID); err != nil {
		return model.CartData{}, err
	} else {
		return model.CartData{ID: time.Now().Day(), UserID: userID, Carts: carts}, nil
	}
}

func (cs *CartService) GetCartByID(id, userID int) (model.Cart, error) {
	if cart, err := cs.cartRepository.GetCartByID(id, userID); err != nil {
		return model.Cart{}, err
	} else {
		return cart, nil
	}
}

func (cs *CartService) UpdateCart(cart model.Cart) (model.Cart, error) { // in cart page

	if cart, err := cs.cartRepository.UpdateCart(cart); err != nil {
		return model.Cart{}, err
	} else {
		return cart, nil
	}
}

func (cs *CartService) RemoveCart(id, userID int) error { // in cart page

	if err := cs.cartRepository.RemoveCart(id, userID); err != nil {
		return err
	}
	return nil

}
