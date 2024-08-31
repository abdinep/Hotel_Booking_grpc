package repository

import (
	"errors"

	"github.com/abdinep/Hotel_Booking_grpc/user_service/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	RegisterUser(user model.User) error
	FindByEmail(email string) (model.User, error)
	FindByID(id uint) (model.User, error)
	CheckUser(userID uint32) bool
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) RegisterUser(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return errors.New("user already exist")
	}
	return nil
}
func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, errors.New("invalid email or password")
	}
	return user, nil
}
func (r *userRepository) FindByID(id uint) (model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}
func (r *userRepository) CheckUser(userID uint32) bool {
	var user model.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return false
	}
	return true
}
