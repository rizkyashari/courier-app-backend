package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	FindById(id uint64) (*entity.SourceOfFund, error)
}

type sourceOfFundRepository struct {
	db *gorm.DB
}

// type SRConfig struct {
// 	DB *gorm.DB
// }

func NewSourceOfFundRepository(c *gorm.DB) SourceOfFundRepository {
	return &sourceOfFundRepository{
		db: c,
	}
}

func (r *sourceOfFundRepository) FindById(id uint64) (*entity.SourceOfFund, error) {
	var sourceOfFund *entity.SourceOfFund

	err := r.db.Where("id = ?", id).Find(&sourceOfFund).Error
	if err != nil {
		return sourceOfFund, err
	}

	return sourceOfFund, nil
}
