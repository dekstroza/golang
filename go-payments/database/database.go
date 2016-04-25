package database

import (
	"github.com/dekstroza/golang/go-payments/models"
	"github.com/jinzhu/gorm"

	//Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

import (
	"fmt"
	"log"
)

//DB is something
var DB *gorm.DB
var err error

//InitDB will initilize database connection
func InitDB(dbHostPtr *string, dbHostPortPtr *int, dbUsernamePtr *string, dbPasswordPtr *string, dbNamePtr *string, dbSchemaPtr *string, dbMaxIdleConnectionsPtr *int, dbMaxConnectionsPtr *int) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *dbHostPtr, *dbHostPortPtr, *dbUsernamePtr, *dbPasswordPtr, *dbNamePtr)
	log.Printf("Connection string is: %s", connectionString)
	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("failed to connect database")
	}
	setDbSchemaString := fmt.Sprintf("SET SEARCH_PATH TO \"%s\"", *dbSchemaPtr)
	log.Printf("%s\n", setDbSchemaString)
	DB.DB().SetMaxIdleConns(*dbMaxIdleConnectionsPtr)
	DB.DB().SetMaxOpenConns(*dbMaxConnectionsPtr)
	DB.Exec(setDbSchemaString)
	DB.AutoMigrate(&models.ApplicationUser{})

}
