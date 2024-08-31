package postgres

import (
	"log"
	"os"

	"github.com/abdinep/Hotel_Booking_grpc/hotel_service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect with postgres......")
	}
	db.AutoMigrate(model.Room{})
	return db
}
