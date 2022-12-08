package repository

import (
	"backend/dto"
	"backend/entity"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository is contract what userRepository can do to db
type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user dto.User) dto.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entity.User
	FindById(id uint64) (*entity.User, error)
	ProfileUser(userID string) entity.User
	Update(user *entity.User) (*entity.User, error)
	AllUsers() []entity.User
	AllUserReferrals() []string
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user dto.User) dto.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser entity.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	db.connection.Debug().Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindById(id uint64) (*entity.User, error) {
	var user *entity.User

	err := db.connection.Where("id =?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *userConnection) FindByEmail(email string) entity.User {
	var user entity.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entity.User {
	var user entity.User
	// db.connection.Preload("Addresses").Preload("Addresses.User").Preload("Shippings").Preload("Shippings.User").Preload("Shippings.Size").Preload("Shippings.Category").Preload("Shippings.Payment").Preload("Shippings.Payment.User").Preload("Shippings.Address").Preload("Shippings.Address.User").Preload("Shippings.AddOn").Find(&user, userID)
	db.connection.Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}

func (db *userConnection) Update(user *entity.User) (*entity.User, error) {
	err := db.connection.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *userConnection) AllUsers() []entity.User {
	var users []entity.User

	db.connection.Preload("User").Preload("Size").Preload("Address").Preload("Address.User").Preload("Payment").Preload("Payment.User").Preload("Category").Preload("AddOn").Find(&users)
	return users
}

func (db *userConnection) AllUserReferrals() []string {
	var referrals []string
	db.connection.Table("users").Select("referral_code").Find(&referrals)
	return referrals
}
