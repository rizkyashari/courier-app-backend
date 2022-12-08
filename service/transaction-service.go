package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"errors"
)

type TransactionService interface {
	GetTransactions(userID uint64, query *dto.TransactionRequestQuery) ([]*entity.Transaction, error)
	TopUp(input *dto.TopUpRequestBody) (*entity.Transaction, error)
	Payment(input *dto.PaymentRequestBody) (*entity.Transaction, error)
	CountTransaction(userID uint64) (int64, error)
}

type transactionService struct {
	transactionRepository  repository.TransactionRepository
	addressRepository      repository.AddressRepository
	userRepository         repository.UserRepository
	paymentRepository      repository.PaymentRepository
	shippingRepository     repository.ShippingRepository
	sourceOfFundRepository repository.SourceOfFundRepository
	promoRepository        repository.PromoRepository
	userPromoRepository    repository.UserPromoRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository, sourceOfFundRepo repository.SourceOfFundRepository, addressRepo repository.AddressRepository, paymentRepo repository.PaymentRepository, shippingRepo repository.ShippingRepository, promoRepo repository.PromoRepository, userPromoRepo repository.UserPromoRepository) TransactionService {
	return &transactionService{
		transactionRepository:  transactionRepo,
		userRepository:         userRepo,
		sourceOfFundRepository: sourceOfFundRepo,
		addressRepository:      addressRepo,
		paymentRepository:      paymentRepo,
		shippingRepository:     shippingRepo,
		promoRepository:        promoRepo,
		userPromoRepository:    userPromoRepo,
	}
}

func (s *transactionService) GetTransactions(userID uint64, query *dto.TransactionRequestQuery) ([]*entity.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll(userID, query)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) TopUp(input *dto.TopUpRequestBody) (*entity.Transaction, error) {
	sourceOfFund, err := s.sourceOfFundRepository.FindById(uint64(input.SourceOfFundID))
	if err != nil {
		return &entity.Transaction{}, err
	}
	if sourceOfFund.ID == 0 {
		return &entity.Transaction{}, errors.New("source of fund is not found")
	}

	user, err := s.userRepository.FindById(uint64(input.User.ID))
	if err != nil {
		return &entity.Transaction{}, err
	}
	if user.ID == 0 {
		return &entity.Transaction{}, errors.New("user is not found")
	}

	transaction := &entity.Transaction{}
	transaction.SourceOfFundID = &sourceOfFund.ID
	transaction.UserID = input.User.ID
	transaction.DestinationID = user.ID
	transaction.Amount = input.Amount
	transaction.Description = "Top Up from " + sourceOfFund.Name
	transaction.Category = "Top Up"

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	user.Balance = user.Balance + input.Amount
	user, err = s.userRepository.Update(user)
	if err != nil {
		return transaction, err
	}

	transaction.SourceOfFund = sourceOfFund
	transaction.User = *input.User
	transaction.User = *user

	return transaction, nil
}

func (s *transactionService) CountTransaction(userID uint64) (int64, error) {
	totalTransactions, err := s.transactionRepository.Count(userID)
	if err != nil {
		return totalTransactions, err
	}

	return totalTransactions, nil
}

func (s *transactionService) Payment(input *dto.PaymentRequestBody) (*entity.Transaction, error) {
	// myAddress, err := s.addressRepository.FindByUserId(uint64(input.User.ID))
	// if err != nil {
	// 	return &entity.Transaction{}, err
	// }
	// if myAddress.ID == 0 {
	// 	return &entity.Transaction{}, errors.New("address is not found")
	// }
	shipping := s.shippingRepository.FindShippingByID(uint64(input.ShippingID))
	userPromo := s.userPromoRepository.FindUserPromoByID(input.UserPromoID)
	promo := s.promoRepository.FindPromoByID(uint64(input.PromoID))

	promo.ID = input.PromoID

	myAddress, err := s.addressRepository.FindById(uint64(shipping.AddressID))
	if err != nil {
		return &entity.Transaction{}, err
	}
	if myAddress.ID == 0 {
		return &entity.Transaction{}, errors.New("shipping id is not found")
	}

	user, _ := s.userRepository.FindById(uint64(input.User.ID))
	if user.Balance < (int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price)) {
		return &entity.Transaction{}, errors.New("insufficient balance")
	}
	// id := strconv.Itoa(int(input.AddressID))
	// id := uint64(input.AddressID)
	// if myAddress.UserID == id {
	// 	return &entity.Transaction{}, errors.New("shipping can't be done to same address")
	// }

	// destinationAddress, err := s.addressRepository.FindById(id)
	// destinationUser, err := s.userRepository.FindById(id)
	// if err != nil {
	// 	return &entity.Transaction{}, err
	// }
	// if destinationUser.ID == 0 {
	// 	return &entity.Transaction{}, errors.New("address is not found")
	// }

	//create transaction for receiver

	transaction := &entity.Transaction{}
	// transaction.UserID = shipping.UserID
	// transaction.DestinationID = shipping.AddressID
	// if promo.ID != 0 {
	// 	transaction.Amount = int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price) - int(promo.MaxDiscount)
	// } else {
	// 	transaction.Amount = int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price)
	// }

	// transaction.Description = input.Description
	// transaction.Category = "Receive Package"

	// transaction, err = s.transactionRepository.Save(transaction)
	// if err != nil {
	// 	return transaction, err
	// }

	// create transaction for sender
	// transaction = &entity.Transaction{}
	transaction.UserID = shipping.UserID
	transaction.DestinationID = shipping.AddressID
	if promo.ID != 0 {
		transaction.Amount = int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price) - int(promo.MaxDiscount)
	} else {
		transaction.Amount = int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price)
	}
	transaction.Description = input.Description
	transaction.Category = "Send Package"

	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	// payment
	// payment, _ := s.paymentRepository.FindById(uint64(input.User.ID))
	payment := &entity.Payment{}
	payment.UserID = input.User.ID
	if userPromo.ID != 0 {
		payment.PromoID = userPromo.PromoID
		payment.PaymentStatus = "Paid with Promo"
		payment.TotalCost = uint64(int(shipping.Category.Price)+int(shipping.Size.Price)+int(shipping.AddOn.Price)) - promo.MaxDiscount
	} else {
		payment.PaymentStatus = "Paid"
		payment.TotalCost = uint64(int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price))
	}
	payment, err = s.paymentRepository.Save(payment)
	if err != nil {
		return transaction, err
	}

	// shipping
	shipping.UserID = input.User.ID
	shipping.PaymentID = payment.ID
	shipping.ShippingStatus = "Paid: delivery is being prepared"
	shipping = s.shippingRepository.UpdateShipping(shipping)

	// userPromo
	userPromo.ID = input.UserPromoID
	userPromo.UserID = input.User.ID
	userPromo.PromoID = promo.ID
	userPromo.Status = 1
	userPromo = s.userPromoRepository.UpdateUserPromo(userPromo)

	totalCost := uint64(payment.TotalCost)
	transaction.SourceOfFundID = &totalCost

	if promo.ID != 0 {
		user.Balance = user.Balance - (int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price)) + int(promo.MaxDiscount)
	} else {
		user.Balance = user.Balance - (int(shipping.Category.Price) + int(shipping.Size.Price) + int(shipping.AddOn.Price))
	}

	user, err = s.userRepository.Update(user)
	if err != nil {
		return transaction, err
	}

	// destinationUser.Balance = destinationUser.Balance + input.Amount
	// _, err = s.userRepository.Update(destinationUser)
	// if err != nil {
	// 	return transaction, err
	// }

	balance := uint64(user.Balance)
	transaction.SourceOfFundID = &balance
	transaction.User = *input.User
	// transaction.User = *destinationUser

	return transaction, nil
}

// func (s *transactionService) Transfer(input *dto.TransferRequestBody) (*entity.Transaction, error) {
// 	myWallet, err := s.walletRepository.FindByUserId(uint64(input.User.ID))
// 	if err != nil {
// 		return &entity.Transaction{}, err
// 	}
// 	if myWallet.ID == 0 {
// 		return &entity.Transaction{}, errors.New("wallet is not found")
// 	}
// 	if myWallet.Balance < input.Amount {
// 		return &entity.Transaction{}, errors.New("insufficient balance")
// 	}
// 	number := strconv.Itoa(input.WalletNumber)
// 	if myWallet.Number == number {
// 		return &entity.Transaction{}, errors.New("transfer can't be done to same wallet")
// 	}

// 	destinationWallet, err := s.walletRepository.FindByNumber(number)
// 	if err != nil {
// 		return &entity.Transaction{}, err
// 	}
// 	if destinationWallet.ID == 0 {
// 		return &entity.Transaction{}, errors.New("wallet is not found")
// 	}

// 	//create transaction for receiver
// 	transaction := &entity.Transaction{}
// 	transaction.UserID = destinationWallet.User.ID
// 	transaction.DestinationID = myWallet.ID
// 	transaction.Amount = input.Amount
// 	transaction.Description = input.Description
// 	transaction.Category = "Receive Money"

// 	transaction, err = s.transactionRepository.Save(transaction)
// 	if err != nil {
// 		return transaction, err
// 	}

// 	// create transaction for sender
// 	transaction = &entity.Transaction{}
// 	transaction.UserID = input.User.ID
// 	transaction.DestinationID = destinationWallet.ID
// 	transaction.Amount = input.Amount
// 	transaction.Description = input.Description
// 	transaction.Category = "Send Money"

// 	transaction, err = s.transactionRepository.Save(transaction)
// 	if err != nil {
// 		return transaction, err
// 	}

// 	myWallet.Balance = myWallet.Balance - input.Amount
// 	myWallet, err = s.walletRepository.Update(myWallet)
// 	if err != nil {
// 		return transaction, err
// 	}

// 	destinationWallet.Balance = destinationWallet.Balance + input.Amount
// 	_, err = s.walletRepository.Update(destinationWallet)
// 	if err != nil {
// 		return transaction, err
// 	}

// 	balance := uint64(myWallet.Balance)
// 	transaction.SourceOfFundID = &balance
// 	transaction.User = *input.User
// 	// transaction.Wallet = *destinationWallet

// 	return transaction, nil
// }
