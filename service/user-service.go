package service

import (
	"backend/dto"
	"backend/entity"
	"backend/repository"
	"log"

	"github.com/mashingan/smapping"
)

type UserService interface {
	Update(user dto.User) dto.User
	GetUser(input *dto.UserRequestParams) (*entity.User, error)
	Profile(userID string) entity.User
	AllUsers() []entity.User
	AllUserReferrals() []string
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) GetUser(input *dto.UserRequestParams) (*entity.User, error) {
	user, err := s.userRepository.FindById(input.UserID)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (service *userService) Update(user dto.User) dto.User {
	userToUpdate := dto.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)

	return updatedUser
}

func (service *userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func (service *userService) AllUsers() []entity.User {
	return service.userRepository.AllUsers()
}

func (service *userService) AllUserReferrals() []string {
	return service.userRepository.AllUserReferrals()
}
