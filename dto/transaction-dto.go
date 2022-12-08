package dto

import (
	"backend/entity"
	"strings"
	"time"
)

type TopUpRequestBody struct {
	Amount         int    `json:"amount" binding:"required,min=10000,max=10000000"`
	SourceOfFundID uint64 `json:"source_of_fund_id" binding:"required"`
	User           *entity.User
}

type PaymentRequestBody struct {
	// Amount int `json:"amount" binding:"required,min=1000,max=50000000"`
	// WalletNumber int    `json:"wallet_number" binding:"required"`
	// AddressID   uint64 `json:"address_id" binding:"required"`
	ShippingID  uint64 `json:"shipping_id" binding:"required"`
	PromoID     uint64 `json:"promo_id"`
	UserPromoID uint64 `json:"user_promo_id"`
	Description string `json:"description"`
	User        *entity.User
	Shipping    *entity.Shipping
	Promo       *entity.Promo
	UserPromo   *entity.UserPromo
}

type TopUpResponse struct {
	ID           uint64 `json:"id"`
	SourceOfFund string `json:"source_of_fund"`
	Amount       int    `json:"amount"`
	// WalletBalance int       `json:"balance"`
	UserBalance int       `json:"balance"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TransactionRequestQuery struct {
	Search string `form:"s"`
	SortBy string `form:"sortBy"`
	Sort   string `form:"sort"`
	Limit  int    `form:"limit"`
	Page   int    `form:"page"`
}

type Destination struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
}

type PaymentResponse struct {
	ID          uint64      `json:"id"`
	Destination Destination `json:"destination"`
	Amount      int         `json:"amount"`
	// WalletBalance int         `json:"balance"`
	// UserBalance int       `json:"balance"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TransactionResponse struct {
	ID           uint64      `json:"id"`
	SourceOfFund string      `json:"source_of_fund"`
	Destination  Destination `json:"destination"`
	Amount       int         `json:"amount"`
	Description  string      `json:"description"`
	Category     string      `json:"category"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
}

func FormatTopUp(transaction *entity.Transaction) TopUpResponse {
	return TopUpResponse{
		ID:           transaction.ID,
		SourceOfFund: transaction.SourceOfFund.Name,
		Amount:       transaction.Amount,
		// WalletBalance: transaction.Wallet.Balance,
		UserBalance: transaction.User.Balance,
		Description: transaction.Description,
		Category:    transaction.Category,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func FormatPayment(transaction *entity.Transaction) PaymentResponse {
	return PaymentResponse{
		ID:          transaction.ID,
		Destination: Destination{Name: transaction.User.Name, ID: transaction.DestinationID},
		Amount:      transaction.Amount,
		// WalletBalance: int(*transaction.SourceOfFundID),
		// UserBalance: transaction.User.Balance,
		Description: transaction.Description,
		Category:    transaction.Category,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func FormatTransaction(transaction *entity.Transaction) TransactionResponse {
	var sourceOfFund string
	if transaction.SourceOfFund != nil {
		sourceOfFund = transaction.SourceOfFund.Name
	}
	return TransactionResponse{
		ID:           transaction.ID,
		SourceOfFund: sourceOfFund,
		Destination:  Destination{Name: transaction.User.Name, ID: transaction.DestinationID},
		// Destination: Destination{Name: transaction.User.Name},
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Category:    transaction.Category,
		CreatedAt:   transaction.CreatedAt,
		UpdatedAt:   transaction.UpdatedAt,
	}
}

func FormatTransactions(transactions []*entity.Transaction) []TransactionResponse {
	formattedTransactions := []TransactionResponse{}
	for _, transaction := range transactions {
		formattedBook := FormatTransaction(transaction)
		formattedTransactions = append(formattedTransactions, formattedBook)
	}
	return formattedTransactions
}

func FormatQuery(query *TransactionRequestQuery) *TransactionRequestQuery {
	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	query.SortBy = strings.ToLower(query.SortBy)
	if query.SortBy == "date" {
		query.SortBy = "updated_at"
	} else if query.SortBy == "to" {
		query.SortBy = "destination_id"
	} else if query.SortBy != "amount" {
		query.SortBy = "updated_at"
	}

	query.Sort = strings.ToUpper(query.Sort)
	if query.Sort != "ASC" {
		query.Sort = "DESC"
	}

	return query
}
