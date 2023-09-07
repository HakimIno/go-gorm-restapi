package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN = "host=db-postgresql-sgp1-29598-do-user-14281553-0.b.db.ondigitalocean.com user=doadmin password=AVNS_sPcl85X3dAS0TgqHBVb dbname=defaultdb port=25060 sslmode=require"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
