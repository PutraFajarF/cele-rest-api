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
		&entities.MasterAuthor{},
		&entities.MasterBook{},
	)

	if err != nil {
		return err
	}

	DB.Migrator().CreateConstraint(&entities.MasterAuthor{}, "MasterBooks")
	DB.Migrator().CreateConstraint(&entities.MasterAuthor{}, "fk_master_authors_master_books")

	log.Println("Database Migration Completed...")
	return nil
}
