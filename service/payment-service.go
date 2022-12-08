package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"fmt"
	"log"

	"github.com/mashingan/smapping"
)

type PaymentService interface {
	Insert(p dto.PaymentCreateDTO) entity.Payment
	Update(p dto.PaymentUpdateDTO) entity.Payment
	Delete(p entity.Payment)
	All() []entity.Payment
	FindByID(paymentID uint64) entity.Payment
	IsAllowedToEdit(userID string, paymentID uint64) bool
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepo,
	}
}

func (service *paymentService) Insert(p dto.PaymentCreateDTO) entity.Payment {
	payment := entity.Payment{}
	err := smapping.FillStruct(&payment, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.paymentRepository.InsertPayment(payment)
	return res
}

func (service *paymentService) Update(p dto.PaymentUpdateDTO) entity.Payment {
	payment := entity.Payment{}
	err := smapping.FillStruct(&payment, smapping.MapFields(&p))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.paymentRepository.UpdatePayment(payment)
	return res
}

func (service *paymentService) Delete(p entity.Payment) {
	service.paymentRepository.DeletePayment(p)
}

func (service *paymentService) All() []entity.Payment {
	return service.paymentRepository.AllPayments()
}

func (service *paymentService) FindByID(paymentID uint64) entity.Payment {
	return service.paymentRepository.FindPaymentByID(paymentID)
}

func (service *paymentService) IsAllowedToEdit(userID string, paymentID uint64) bool {
	b := service.paymentRepository.FindPaymentByID(paymentID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
