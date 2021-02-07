package database

import (
	"log"

	"github.com/glup3/jetti-ui/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect connects to postgres database
func Connect() *gorm.DB {
	dsn := config.Get("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Couldn't connect to database")
	}

	return db
}
