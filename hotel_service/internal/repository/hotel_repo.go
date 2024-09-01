package repository

import (
	"errors"
	"fmt"

	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/model"
	"gorm.io/gorm"
)

type RoomRepository interface {
	AddRoom(room model.Room) error
	GetRoom(roomNumber string) (model.Room, error)
	UpdateRoom(room model.Room) error
	DeleteRoom(roomNumber string) error
	GetRooms() ([]model.Room, error)
	RoomAvailable(roomID string) bool
}

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) AddRoom(room model.Room) error {
	if err := r.db.Create(&room).Error; err != nil {
		return errors.New("failed to add room")
	}
	return nil
}

func (r *roomRepository) GetRoom(roomNumber string) (model.Room, error) {
	var room model.Room
	fmt.Println(roomNumber)
	if err := r.db.First(&room, "room_number = ?", roomNumber).Error; err != nil {
		return model.Room{}, errors.New("room not found")
	}
	return room, nil
}

func (r *roomRepository) UpdateRoom(room model.Room) error {
	if err := r.db.Where("room_number = ?", room.RoomNumber).Save(&room).Error; err != nil {
		return errors.New("failed to update room")
	}
	return nil
}

func (r *roomRepository) DeleteRoom(roomNumber string) error {
	if err := r.db.Delete(&model.Room{}, "room_number = ?", roomNumber).Error; err != nil {
		return errors.New("failed to delete room")
	}
	return nil
}

func (r *roomRepository) GetRooms() ([]model.Room, error) {
	var rooms []model.Room
	if err := r.db.Find(&rooms).Error; err != nil {
		return nil, errors.New("failed to get rooms")
	}
	return rooms, nil
}
func (r *roomRepository) RoomAvailable(roomID string) bool {
	var room model.Room
	if err := r.db.Where("availability = ?", true).First(&room, "room_number = ?", roomID).Error; err != nil {
		return false
	}
	return true
}
