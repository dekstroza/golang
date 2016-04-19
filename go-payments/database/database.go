package database

import (
	"github.com/dekstroza/golang/go-payments/models"
	"github.com/jinzhu/gorm"

	//Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DB is something
var DB *gorm.DB
var err error

//InitDB will initilize database connection
func InitDB() {
	DB, err = gorm.Open("postgres", "host=172.18.0.2 user=postgres password=mysecretpassword dbname=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.ApplicationUser{})

}
