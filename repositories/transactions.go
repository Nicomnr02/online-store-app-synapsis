package repositories

import (
	"online_app_store/model"

	"gorm.io/gorm"
)

type TransactionRepositoryInterface interface {
	GetAllTransactionsByUserID(userID int) ([]model.Transaction, error)
	CreateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error)
	UpdateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error)
	DeleteTransaction(transactionReq model.TransactionRequest) error
}

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (tr *TransactionRepository) GetAllTransactionsByUserID(userID int) ([]model.Transaction, error) {
	var transactions []model.Transaction

	if res := tr.db.Table("transactions").Joins("INNER JOIN users on transactions.user_id = users.id").Find(&transactions, "transactions.user_id = ?", userID); res.Error != nil {
		return []model.Transaction{}, res.Error
	}

	return transactions, nil
}

func (tr *TransactionRepository) CreateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error) {
	newTransaction := model.Transaction{UserID: transactionReq.UserID, CartID: transactionReq.CartID, Status: "not_paid"}
	res := tr.db.Save(&newTransaction)
	if res.Error != nil {
		return model.Transaction{}, res.Error
	}

	return newTransaction, nil
}

func (tr *TransactionRepository) UpdateTransaction(transactionReq model.TransactionRequest) (model.Transaction, error) {
	newTransaction := model.Transaction{ID: transactionReq.ID, UserID: transactionReq.UserID, CartID: transactionReq.CartID, Status: "paid"}
	if err := tr.db.Model(&newTransaction).Updates(&newTransaction).Error; err != nil {
		return model.Transaction{}, err
	}
	return newTransaction, nil
}

func (tr *TransactionRepository) DeleteTransaction(transactionReq model.TransactionRequest) error {
	if res := tr.db.Joins("INNER JOIN users on transactions.user_id = users.id").Where("id = ? AND transactions.user_id = ?", transactionReq.ID, transactionReq.UserID).Unscoped().Delete(&model.Transaction{}); res.Error != nil {
		return res.Error
	}
	return nil
}
