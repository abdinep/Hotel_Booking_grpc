package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpcphandle "github.com/abdinep/Hotel_Booking_grpc/user_service/internal/delivery/grpc"
	httphandle "github.com/abdinep/Hotel_Booking_grpc/user_service/internal/delivery/http"
	"github.com/abdinep/Hotel_Booking_grpc/user_service/internal/repository"
	"github.com/abdinep/Hotel_Booking_grpc/user_service/internal/service"
	middlewaregrpc "github.com/abdinep/Hotel_Booking_grpc/user_service/pkg/middleware-grpc"
	"github.com/abdinep/Hotel_Booking_grpc/user_service/pkg/postgres"
	user "github.com/abdinep/Hotel_Booking_grpc/user_service/proto"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	logs "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
	err := godotenv.Load("/home/abdin/Broto/Brocamp/week-20 (microservice)/Hotel_booking_grpc/user_service/.env")
	if err != nil {
		log.Println("No .env file found")
	}
	logs.SetOutput(os.Stdout)

	logs.SetLevel(logs.InfoLevel)

	logs.SetFormatter(&logs.JSONFormatter{})
}
func main() {

	log.Println("Starting application...")
	conn, err := grpc.NewClient("localhost:50042", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	db := postgres.InitDatabase()
	repo := repository.NewUserRepository(db)
	userService := service.NewUserService(repo, conn)
	grpchandler := grpcphandle.NewUserHandler(userService)
	httpHandler := httphandle.NewUserHandler(userService)

	secureMethods := map[string]bool{
		"/proto.UserService/UserGetInfo":   true,
		"/proto.UserService/GetHotelRooms": true,
	}
	go func() {
		r := gin.Default()
		httpHandler.UserRouters(r)
		if err := r.Run(":8081"); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
		fmt.Println("gin server is running")
	}()

	go func() {
		list, err := net.Listen("tcp", ":50041")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer(grpc.UnaryInterceptor(middlewaregrpc.AuthInterceptor(secureMethods)))
		user.RegisterUserServiceServer(s, grpchandler)

		log.Printf("Server listening on port 50041")
		if err := s.Serve(list); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down servers...")
}
