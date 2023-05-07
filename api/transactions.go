package api

import (
	"encoding/json"
	"net/http"
	"online_app_store/model"
	"online_app_store/services"
)

type TransactionAPIInterface interface {
	GetAllTransactionsByUserID(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
	CreateTransactions(w http.ResponseWriter, r *http.Request)
	UpdateTransaction(w http.ResponseWriter, r *http.Request)
	DeleteTransaction(w http.ResponseWriter, r *http.Request)
}

type TransactionAPI struct {
	transService services.TransactionServiceInterface
}

func NewTransactionAPI(transService services.TransactionServiceInterface) *TransactionAPI {
	return &TransactionAPI{transService}
}

func (ta *TransactionAPI) GetAllTransactionsByUserID(w http.ResponseWriter, r *http.Request) {
	userID := int(r.Context().Value("user_id").(float64))

	if transactions, err := ta.transService.GetAllTransactionsByUserID(userID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success get all transactions by user id", Data: transactions})
	}
}

func (ta *TransactionAPI) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := int(r.Context().Value("user_id").(float64))
	transactionRequest.UserID = userID

	if transactions, err := ta.transService.CreateTransaction(transactionRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success create a transaction ", Data: transactions})
	}
}

func (ta *TransactionAPI) CreateTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionRequest []model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := int(r.Context().Value("user_id").(float64))

	for idx, tr := range transactionRequest {
		tr.UserID = userID
		transactionRequest[idx] = tr
	}

	if transactions, err := ta.transService.CreateTransactions(transactionRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success create transactions ", Data: transactions})
	}
}

func (ta *TransactionAPI) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := int(r.Context().Value("user_id").(float64))
	transactionRequest.UserID = userID

	if transactions, err := ta.transService.UpdateTransaction(transactionRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success update transaction", Data: transactions})
	}
}

func (ta *TransactionAPI) UpdateTransactions(w http.ResponseWriter, r *http.Request) {
	var transactionRequest []model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := int(r.Context().Value("user_id").(float64))

	for idx, tr := range transactionRequest {
		tr.UserID = userID
		transactionRequest[idx] = tr
	}

	if transactions, err := ta.transService.UpdateTransactions(transactionRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success update choosen transaction", Data: transactions})
	}
}

func (ta *TransactionAPI) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var transactionRequest model.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	}

	userID := int(r.Context().Value("user_id").(float64))
	transactionRequest.UserID = userID

	if err := ta.transService.DeleteTransaction(transactionRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&model.Response{Message: err.Error()})
		return
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&model.Response{Message: "Success delete transaction by user id"})
	}
}
