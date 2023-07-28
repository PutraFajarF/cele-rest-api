package config

import (
	"fmt"
	"log"
	"project-rest-api/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() (*gorm.DB, error) {
	dataConn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		AppConfig.Host,
		AppConfig.Username,
		AppConfig.Password,
		AppConfig.DatabaseName,
		AppConfig.Port,
	)

	DB, err = gorm.Open(postgres.Open(dataConn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	return DB, nil
}

func Migrate() error {
	err := DB.AutoMigrate(
		&entities.User{},
	)

	if err != nil {
		return err
	}

	log.Println("Database Migration Completed...")
	return nil
}
