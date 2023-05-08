package services

import (
	"errors"
	"fmt"
	"online_app_store/model"
	"online_app_store/repositories"
)

type TransactionServiceInterface interface {
	GetAllTransactionsByUserID(userID int) ([]model.Transaction, error)
	CreateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error)
	CreateTransactions(transactionsReq []model.TransactionRequest) ([]model.Transaction, error)
	UpdateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error)
	UpdateTransactions(transactionsReq []model.TransactionRequest) ([]model.Transaction, error)
	DeleteTransaction(transactionReq model.TransactionRequest) error
}

type TransactionService struct {
	userRepository    repositories.UserRepositoryInterface
	productRepository repositories.ProductRepositoryInterface
	cartRepository    repositories.CartRepositoryInterface
	transRepository   repositories.TransactionRepositoryInterface
}

func NewTransactionService(userRepository repositories.UserRepositoryInterface, productRepository repositories.ProductRepositoryInterface, cartRepository repositories.CartRepositoryInterface, transRepository repositories.TransactionRepositoryInterface) TransactionServiceInterface {
	return &TransactionService{userRepository, productRepository, cartRepository, transRepository}
}

func (ts *TransactionService) GetAllTransactionsByUserID(userID int) ([]model.Transaction, error) {
	if transactions, err := ts.transRepository.GetAllTransactionsByUserID(userID); err != nil {
		return []model.Transaction{}, err
	} else {
		return transactions, nil
	}
}

func (ts *TransactionService) CreateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error) {
	if transaction, err := ts.transRepository.CreateTransaction(transactionReq); err != nil {
		return model.Transaction{}, err
	} else {
		return transaction, nil
	}
}

func (ts *TransactionService) CreateTransactions(transactionsReq []model.TransactionRequest) ([]model.Transaction, error) {
	carts, err := ts.cartRepository.GetAllCartsByUserID(transactionsReq[0].UserID)
	if err != nil {
		return []model.Transaction{}, err
	}

	var transactions []model.Transaction
	for _, cart := range carts {
		newTransaction := model.TransactionRequest{UserID: cart.UserID, CartID: cart.ID}
		if transaction, err := ts.transRepository.CreateTransaction(newTransaction); err != nil {
			return []model.Transaction{}, err
		} else {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

func (ts *TransactionService) UpdateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error) {

	//! getCash
	userData, err := ts.userRepository.GetUserByID(transactionReq.UserID)
	if err != nil {
		return model.Transaction{}, err
	}
	cash := userData.Cash

	//! getCart
	allTrans, err := ts.transRepository.GetAllTransactionsByUserID(transactionReq.UserID)
	if err != nil {
		return model.Transaction{}, err
	}

	var cartData model.Cart
	for _, transac := range allTrans {
		if transac.ID == transactionReq.ID {
			choosenCart, err := ts.cartRepository.GetCartByID(transac.CartID, transac.UserID)
			if err != nil {
				return model.Transaction{}, err
			}
			cartData = choosenCart
		}
	}

	cash -= cartData.Price

	if cash < 0 {
		return model.Transaction{}, errors.New("not enough money")
	}

	if cash == 0 {
		userData.Cash = -1
	} else {
		userData.Cash = cash
	}

	//! update user (money)
	if _, err := ts.userRepository.UpdateUser(userData); err != nil {
		return model.Transaction{}, err
	}

	//! delete paid cart
	cartsData, err := ts.cartRepository.GetAllCartsByUserID(userData.ID)
	if err != nil {
		return model.Transaction{}, err
	}

	for _, cart := range cartsData {
		if cartData.ID == cart.ID {
			if err := ts.cartRepository.RemoveCart(cart.ID, cart.UserID); err != nil {
				return model.Transaction{}, err
			}
		}

	}

	//! update transaction
	newTransaction, err := ts.transRepository.UpdateTransaction(transactionReq)
	if err != nil {
		return model.Transaction{}, err
	}

	//! update product stock
	for _, cart := range cartsData {
		prodID := cart.ProductID
		product, err := ts.productRepository.GetProductByID(prodID)
		if err != nil {
			return model.Transaction{}, err
		}

		product.Stock -= cart.Quantity

		updatedProduct, err := ts.productRepository.UpdateProduct(product)
		if err != nil {
			return model.Transaction{}, err
		}

		fmt.Println("updated stock : ", updatedProduct)

	}

	return newTransaction, nil
}

func (ts *TransactionService) UpdateTransactions(transactionsReq []model.TransactionRequest) ([]model.Transaction, error) {
	//! getCash
	userData, err := ts.userRepository.GetUserByID(transactionsReq[0].UserID)
	if err != nil {
		return []model.Transaction{}, err
	}
	cash := userData.Cash

	//! getCart
	var choosenCarts []model.Cart
	transactions, err := ts.transRepository.GetAllTransactionsByUserID(userData.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	for _, existedTransaction := range transactions {
		for _, reqTransaction := range transactionsReq {
			if reqTransaction.ID == existedTransaction.ID {

				cart, err := ts.cartRepository.GetCartByID(existedTransaction.CartID, existedTransaction.UserID)
				if err != nil {
					return []model.Transaction{}, err
				}
				choosenCarts = append(choosenCarts, cart)
			}
		}
	}

	//! do transaction
	for _, cart := range choosenCarts {
		cash -= cart.Price
	}

	if cash < 0 {
		return []model.Transaction{}, errors.New("not enough money")
	}

	if cash == 0 {
		userData.Cash = -1
	} else {
		userData.Cash = cash
	}

	//! update user (money) if >= 0
	userData.Cash = cash
	if _, err := ts.userRepository.UpdateUser(userData); err != nil {
		return []model.Transaction{}, err
	}

	//! delete paid cart
	allCarts, err := ts.cartRepository.GetAllCartsByUserID(userData.ID)
	if err != nil {
		return []model.Transaction{}, err
	}

	for _, existedCart := range allCarts {
		for _, choosenCart := range choosenCarts {
			if existedCart.ID == choosenCart.ID {
				err := ts.cartRepository.RemoveCart(existedCart.ID, existedCart.UserID)
				if err != nil {
					return []model.Transaction{}, err
				}
			}
		}

	}

	//! update paid transaction
	var updatedTransactions []model.Transaction
	for _, tr := range transactionsReq {
		updatedTrans, err := ts.transRepository.UpdateTransaction(tr)
		if err != nil {
			return []model.Transaction{}, err
		}
		updatedTransactions = append(updatedTransactions, updatedTrans)
	}

	//! update paid product stock
	for _, cart := range choosenCarts {
		prodID := cart.ProductID
		product, err := ts.productRepository.GetProductByID(prodID)
		if err != nil {
			return []model.Transaction{}, err
		}

		product.Stock -= cart.Quantity

		updatedProduct, err := ts.productRepository.UpdateProduct(product)
		if err != nil {
			return []model.Transaction{}, err
		}

		fmt.Println("updated stock : ", updatedProduct)

	}

	return updatedTransactions, nil
}

func (ts *TransactionService) DeleteTransaction(transactionReq model.TransactionRequest) error {
	if err := ts.transRepository.DeleteTransaction(transactionReq); err != nil {
		return err
	}

	return nil
}
