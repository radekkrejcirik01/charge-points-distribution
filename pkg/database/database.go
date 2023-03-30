package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbhost     = "localhost"
	dbport     = "5432"
	dbuser     = "user"
	dbpassword = "userpassword"
	dbname     = "chargepoints"
)

// DB is connected Postgres DB
var DB *gorm.DB

// Connect to Postgres server
func Connect() {
	dsn := "host=host.docker.internal user=user password=userpassword dbname=chargepoints port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

// GetConfig for debuging
func GetConfig() (string, string, string, string, string) {
	return dbhost, dbport, dbuser, dbpassword, dbname
}
