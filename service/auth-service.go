package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRep repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRep,
	}
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service *authService) CreateUser(user dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	if userToCreate.Name == "Admin" && userToCreate.Email == "admin@gmail.com" {
		userToCreate.Role = "Admin"
	} else {
		userToCreate.Role = "User"
		name := strings.ToUpper(userToCreate.Name)
		rand.Seed(time.Now().Unix())
		str := "1234567890"
		shuff := []rune(str)
		// Shuffling the string
		rand.Shuffle(len(shuff), func(i, j int) {
			shuff[i], shuff[j] = shuff[j], shuff[i]
		})
		// Displaying the random string
		fmt.Println(string(shuff))
		userToCreate.ReferralCode = name[0:4] + string(shuff)[0:4]

		var users []string = service.userRepository.AllUserReferrals()
		s := users
		if contains(s, user.ReferralCode) && user.ReferralCode != "" {
			fmt.Println("Referral Code is Used")
			// payment := &entity.Payment{}
			userToCreate.Balance = userToCreate.Balance + 50000
		} else {
			fmt.Println("Referral Code is not match")
		}
	}

	res := service.userRepository.InsertUser(userToCreate)

	return res
}

func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
