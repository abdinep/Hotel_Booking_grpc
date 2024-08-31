package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/abdinep/Hotel_Booking_grpc/user_service/internal/repository"
	"github.com/abdinep/Hotel_Booking_grpc/user_service/model"
	middleware "github.com/abdinep/Hotel_Booking_grpc/user_service/pkg/middlerware-http"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(name, email, password, mobile string) error
	LoginUser(email, password string) (string, error)
	UserGetInfo(id uint) (model.User, error)
	// GetHotelRooms() ([]*hotel_service.Room, error)
	UserExists(userID uint32) bool
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
func (s *userService) RegisterUser(name, email, password, mobile string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Mobile:   mobile,
	}

	err = s.repo.RegisterUser(user)
	return err
}
func (s *userService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &middleware.Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		return "", err
	}
	fmt.Println("jwt", user.ID)
	return tokenString, nil
}
func (s *userService) UserGetInfo(id uint) (model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (s *userService) UserExists(userID uint32) bool {
	logrus.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("Checking user")

	check := s.repo.CheckUser(userID)
	if !check {
		logrus.WithFields(logrus.Fields{
			"userID": userID,
			"status": check,
		}).Error("User not found")
		return false
	}

	logrus.WithFields(logrus.Fields{
		"userID": userID,
	}).Info("User found")
	return true
}
