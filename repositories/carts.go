package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type CartRepositoryInterface interface {
	CreateCart(cart model.Cart) (model.Cart, error)
	GetAllCartsByUserID(userID int) ([]model.Cart, error)
	GetCartByID(id, userID int) (model.Cart, error)
	UpdateCart(cart model.Cart) (model.Cart, error)
	RemoveCart(id, userID int) error
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db}
}

func (cr *CartRepository) CreateCart(cart model.Cart) (model.Cart, error) {
	var newCart = cart

	if err := cr.db.Save(&newCart).Error; err != nil {
		return model.Cart{}, err
	}

	return newCart, nil
}

func (cr *CartRepository) GetAllCartsByUserID(userID int) ([]model.Cart, error) {
	var carts []model.Cart

	if res := cr.db.Table("carts").Joins("INNER JOIN users on carts.user_id = users.id").Find(&carts, "carts.user_id = ?", userID); res.Error != nil {
		return []model.Cart{}, res.Error
	}

	return carts, nil
}

func (cr *CartRepository) GetCartByID(id, userID int) (model.Cart, error) {
	var cart model.Cart

	if res := cr.db.Table("carts").Joins("INNER JOIN users on carts.user_id = users.id").Where("carts.id = ? AND carts.user_id = ?", id, userID).Find(&cart); res.Error != nil {
		return model.Cart{}, res.Error
	}
	return cart, nil
}

func (cr *CartRepository) UpdateCart(cart model.Cart) (model.Cart, error) {
	var newCart = cart
	if err := cr.db.Model(&cart).Updates(&newCart).Error; err != nil {
		return model.Cart{}, err
	}

	return newCart, nil
}

func (cr *CartRepository) RemoveCart(id, userID int) error {

	if res := cr.db.Joins("INNER JOIN users on carts.user_id = users.id").Where("id = ? AND carts.user_id = ?", id, userID).Unscoped().Delete(&model.Cart{}); res.Error != nil {
		return res.Error
	}
	return nil
}
