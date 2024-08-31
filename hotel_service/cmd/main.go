package main

import (
	"fmt"
	"log"
	"net"

	grpchandler "github.com/abdinep/Hotel_Booking_grpc/hotel_service/internal/delivery/grpc"
	httphandler "github.com/abdinep/Hotel_Booking_grpc/hotel_service/internal/delivery/http"
	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/internal/repository"
	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/internal/service"
	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/pkg/postgres"
	pb "github.com/abdinep/Hotel_Booking_grpc/hotel_service/proto"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load("/home/abdin/Broto/Brocamp/week-20 (microservice)/Hotel_booking_grpc/hotel_service/.env")
	if err != nil {
		log.Println("No .env file found")
	}
	db := postgres.InitDatabase()
	repo := repository.NewRoomRepository(db)
	hotelService := service.NewHotelService(repo)
	r := gin.Default()
	httpHandler := httphandler.NewHotelHandler(hotelService)
	httpHandler.HotelRouters(r)
	grpcServer := grpc.NewServer()
	grpcHandler := grpchandler.NewHotelHandler(hotelService)
	pb.RegisterHotelServiceServer(grpcServer, grpcHandler)

	go func() {
		if err := r.Run(":8082"); err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
		fmt.Println("gin server is running")
	}()

	lis, err := net.Listen("tcp", ":50042")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is running on port 50042")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
