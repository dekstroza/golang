package database

import (
	"github.com/dekstroza/golang/go-payments/models"
	"github.com/dekstroza/golang/go-payments/utils"
	"github.com/jinzhu/gorm"

	//Import postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

import "fmt"

//DB points to gorm DB structure
var DB *gorm.DB
var err error

//InitDB will initilize database connection
func InitDB(cmdLineFlags *utils.CmdLineFlags) *gorm.DB {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *cmdLineFlags.DbHost, *cmdLineFlags.DbPort, *cmdLineFlags.DbUsername, *cmdLineFlags.DbPassword, *cmdLineFlags.DbName)

	DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		panic("Failed to connect database")
	}
	setDbSchemaString := fmt.Sprintf("SET SEARCH_PATH TO \"%s\"", *cmdLineFlags.DbSchema)

	DB.DB().SetMaxIdleConns(*cmdLineFlags.MaxIdleConnections)
	DB.DB().SetMaxOpenConns(*cmdLineFlags.MaxConnections)
	DB.Exec(setDbSchemaString)
	DB.AutoMigrate(&models.ApplicationUser{})
	DB.Debug()
	return DB

}
