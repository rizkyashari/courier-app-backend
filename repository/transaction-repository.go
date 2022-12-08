package repository

import (
	"backend/dto"
	"backend/entity"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll(userID uint64, query *dto.TransactionRequestQuery) ([]*entity.Transaction, error)
	Count(userID uint64) (int64, error)
	Save(transaction *entity.Transaction) (*entity.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(c *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: c,
	}
}

func (r *transactionRepository) FindAll(userID uint64, query *dto.TransactionRequestQuery) ([]*entity.Transaction, error) {
	var transactions []*entity.Transaction

	offset := (query.Page - 1) * query.Limit
	orderBy := query.SortBy + " " + query.Sort
	queryBuider := r.db.Limit(query.Limit).Offset(offset).Order(orderBy)
	err := queryBuider.Debug().Where("user_id = ?", userID).Where("description ILIKE ?", "%"+query.Search+"%").Preload("SourceOfFund").
		Preload("User").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) Count(userID uint64) (int64, error) {
	var total int64
	db := r.db.Model(&entity.Transaction{}).Where("user_id = ?", userID).Count(&total)

	if db.Error != nil {
		return 0, db.Error
	}

	return total, nil
}

func (r *transactionRepository) Save(transaction *entity.Transaction) (*entity.Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
