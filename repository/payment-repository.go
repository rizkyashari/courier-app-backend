package repository

import (
	"backend/entity"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	InsertPayment(p entity.Payment) entity.Payment
	UpdatePayment(p entity.Payment) entity.Payment
	DeletePayment(p entity.Payment)
	AllPayments() []entity.Payment
	FindPaymentByID(paymentID uint64) entity.Payment
	Update(payment *entity.Payment) (*entity.Payment, error)
	FindById(id uint64) (*entity.Payment, error)
	Save(payment *entity.Payment) (*entity.Payment, error)
}

type paymentConnection struct {
	connection *gorm.DB
}

func NewPaymentRepository(dbConn *gorm.DB) PaymentRepository {
	return &paymentConnection{
		connection: dbConn,
	}
}

func (db *paymentConnection) InsertPayment(p entity.Payment) entity.Payment {
	db.connection.Save(&p)
	db.connection.Preload("User").Preload("Promo").Find(&p)
	return p
}

func (db *paymentConnection) UpdatePayment(p entity.Payment) entity.Payment {
	db.connection.Save(&p)
	db.connection.Preload("User").Preload("Promo").Find(&p)
	return p
}

func (db *paymentConnection) DeletePayment(p entity.Payment) {
	db.connection.Delete(&p)
}

func (db *paymentConnection) FindPaymentByID(paymentID uint64) entity.Payment {
	var payment entity.Payment
	db.connection.Preload("User").Preload("Promo").Find(&payment, paymentID)
	return payment
}

func (db *paymentConnection) AllPayments() []entity.Payment {
	var payments []entity.Payment
	db.connection.Preload("User").Preload("Promo").Find(&payments)
	return payments
}

func (db *paymentConnection) Update(payment *entity.Payment) (*entity.Payment, error) {
	err := db.connection.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (db *paymentConnection) FindById(id uint64) (*entity.Payment, error) {
	var payment *entity.Payment

	err := db.connection.Where("id =?", id).Find(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (db *paymentConnection) Save(payment *entity.Payment) (*entity.Payment, error) {
	err := db.connection.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}
