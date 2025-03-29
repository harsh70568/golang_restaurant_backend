package db

import (
	"fmt"
	"golang_restaurant_backend/models"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ServerPort string
var SecretKey, RefreshSecretKey []byte

func ConnectDB() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	hostName := viper.GetString("DB_HOST")
	userName := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")
	ServerPort = viper.GetString("SERVER_PORT")
	SecretKey = []byte(viper.GetString("JWT_SECRET"))
	RefreshSecretKey = []byte(viper.GetString("JWT_REFRESH_SECRET"))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", hostName, userName, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	DB = db
	log.Println("Connected to database succesfully...")
}

func MigrateDatabase() {
	DB.AutoMigrate(&models.Food{})
	DB.AutoMigrate(&models.Menu{})
	DB.AutoMigrate(&models.Order{})
	DB.AutoMigrate(&models.OrderItem{})
	DB.AutoMigrate(&models.Table{})
	DB.AutoMigrate(&models.User{})
}
