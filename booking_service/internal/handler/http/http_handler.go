package http

import (
	"net/http"

	"github.com/abdinep/Hotel_Booking_grpc/booking_service/internal/model"
	"github.com/abdinep/Hotel_Booking_grpc/booking_service/internal/service"
	http_middleware "github.com/abdinep/Hotel_Booking_grpc/booking_service/pkg/middlerware_http"
	"github.com/gin-gonic/gin"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(bookingService service.BookingService) *BookingHandler {
	return &BookingHandler{bookingService: bookingService}
}

func (h *BookingHandler) BookingRouters(r *gin.Engine) {
	r.POST("/book", http_middleware.AuthMiddleware(), h.CreateBooking)
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var req model.Booking
	userId := c.GetUint("userID")
	// Bind JSON request to the struct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	booking, err := h.bookingService.CreateBooking(uint32(userId), req.RoomID, req.CheckIn, req.CheckOut, req.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return response
	c.JSON(http.StatusOK, gin.H{
		"booking_id":     booking.OrderId,
		"razor_order_id": booking.RazorOrderId,
		"status":         booking.Status,
	})
}
