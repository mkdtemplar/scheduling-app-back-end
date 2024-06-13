package db

import (
	"fmt"
	"log"
	"scheduling-app-back-end/internal/models"
	"scheduling-app-back-end/internal/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn = &PostgresDB{}

func ConnectToPostgres() {
	dbConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	dbConn, err := gorm.Open(postgres.Open(dbConfig.DbSource), &gorm.Config{})

	if err != nil {
		fmt.Println("Cannot connect to  database")
		log.Fatal("This is the error:", err)
	} else {
		fmt.Println("We are connected to the database")
	}

	err = dbConn.AutoMigrate(&models.Schedule{}, &models.Positions{}, &models.Users{}, &models.Shifts{},
		&models.Admin{}, &models.AnnualLeave{})
	if err != nil {
		log.Fatal(err)
	}

	Conn.DB = dbConn
}

func GetDb() *gorm.DB {
	return Conn.DB
}
