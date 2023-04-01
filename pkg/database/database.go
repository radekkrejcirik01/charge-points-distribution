package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbhost     = "host.docker.internal"
	dbport     = "5432"
	dbuser     = "user"
	dbpassword = "userpassword"
	dbname     = "distribution"
)

var DB *gorm.DB

// Connect to Postgres server
func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbhost,
		dbuser,
		dbpassword,
		dbname,
		dbport,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = database
	log.Println("succesfully connected to database")
}
