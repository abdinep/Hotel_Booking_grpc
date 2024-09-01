package grpc

import (
	"context"

	"github.com/abdinep/Hotel_Booking_grpc/booking_service/internal/service"
	pb "github.com/abdinep/Hotel_Booking_grpc/booking_service/proto"
)

type BookingHandler struct {
	service service.BookingService
	pb.UnimplementedBookingServiceServer
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{service: service}
}

func (b *BookingHandler) BookingComplete(ctx context.Context, req *pb.BookingCompleteRequest) (*pb.BookingCompleteResponse, error) {
	status, err := b.service.BookingComplete(req.OrderId, req.Status)
	return &pb.BookingCompleteResponse{Status: status}, err

}

// func (h *BookingHandler) CreateBooking(ctx context.Context, req *pb.CreateBookingRequest) (*pb.CreateBookingResponse, error) {
// 	checkIn := time.Unix(req.CheckIn.Seconds, int64(req.CheckIn.Nanos))
// 	checkOut := time.Unix(req.CheckOut.Seconds, int64(req.CheckOut.Nanos))
// 	booking, err := h.service.CreateBooking(ctx, req.UserId, req.RoomId, checkIn, checkOut, req.Amount)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.CreateBookingResponse{
// 		BookingId: booking.ID,
// 		Status:    booking.Status,
// 	}, nil
// }
