package configs

import (
	"fmt"
	"krsan/database"
	"krsan/models"
	"krsan/models/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db_name = "test"
var db_port = "3306"
var db_user = "root"
var db_password = ""
var db_host = "localhost"

func BootDatabase() {
	if dbPortEnv := os.Getenv("DB_PORT"); dbPortEnv != "" {
		db_port = dbPortEnv
	}

	if dbNameEnv := os.Getenv("DB_NAME"); dbNameEnv != "" {
		db_name = dbNameEnv
	}

	if dbUserEnv := os.Getenv("DB_USER"); dbUserEnv != "" {
		db_user = dbUserEnv
	}

	if dbPasswordEnv := os.Getenv("DB_PASSWORD"); dbPasswordEnv != "" {
		db_password = dbPasswordEnv
	}

	if dbHostEnv := os.Getenv("DB_HOST"); dbHostEnv != "" {
		db_host = dbHostEnv
	}
}

func LoadConfig() models.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	// serverPort := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := models.Config{
		// ServerPort: serverPort,
		Database: models.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
		},
	}
	return config
}

func ConnectDatabase() {

	var errConnection error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_password, db_host, db_port, db_name)
	database.DB, errConnection = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConnection != nil {
		panic("Can't connect to database")
	} else {
		log.Println("Database connected")
	}

}

func RunMigration() {
	err := database.DB.AutoMigrate(
		entity.Mahasiswa{},
		entity.JadwalKuliah{},
		entity.Dosen{},
		entity.Krs{},
		entity.MataKuliah{},
	)

	if err != nil {
		fmt.Println(err)
		log.Println("Failed to migrate schema")
	} else {
		log.Println("schemas migrated")
	}
}
