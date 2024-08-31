package service

import (
	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/model"
	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/internal/repository"
)

type HotelService interface {
	AddRoom(room model.Room) error
	GetRoom(roomNumber string) (model.Room, error)
	UpdateRoom(room model.Room) error
	DeleteRoom(roomNumber string) error
	GetRooms() ([]model.Room, error)
	RoomAvailable(roomID string) bool
}

type hotelService struct {
	repo repository.RoomRepository
}

func NewHotelService(repo repository.RoomRepository) HotelService {
	return &hotelService{repo: repo}
}

func (s *hotelService) AddRoom(room model.Room) error {
	return s.repo.AddRoom(room)
}

func (s *hotelService) GetRoom(roomNumber string) (model.Room, error) {
	return s.repo.GetRoom(roomNumber)
}

func (s *hotelService) UpdateRoom(room model.Room) error {
	return s.repo.UpdateRoom(room)
}

func (s *hotelService) DeleteRoom(roomNumber string) error {
	return s.repo.DeleteRoom(roomNumber)
}

func (s *hotelService) GetRooms() ([]model.Room, error) { // Implement this function
	return s.repo.GetRooms()
}
func (s *hotelService) RoomAvailable(roomID string) bool {
	return s.repo.RoomAvailable(roomID)
}
