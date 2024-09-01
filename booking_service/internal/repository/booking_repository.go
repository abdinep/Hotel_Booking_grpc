package repository

import (
	"errors"

	"github.com/abdinep/Hotel_Booking_grpc/booking_service/internal/model"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
	UpdateBooking(booking *model.Booking) error
	GetBookingByID(id uint) (*model.Booking, error)
	BookingComplete(razorId string, status string) (string, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) error {
	if err := r.db.Create(booking).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) UpdateBooking(booking *model.Booking) error {
	if err := r.db.Save(booking).Error; err != nil {
		return err
	}
	return nil
}

func (r *bookingRepository) GetBookingByID(id uint) (*model.Booking, error) {
	var booking model.Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &booking, nil
}
func (b *bookingRepository) BookingComplete(razorId string, status string) (string, error) {
	var booking model.Booking
	if err := b.db.Where("razor_order_id = ?", razorId).First(&booking).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "Booking not found", err
		}
		return "", err
	}
	booking.Status = status
	if err := b.db.Save(&booking).Error; err != nil {
		return "", err
	}
	return "success", nil
}
