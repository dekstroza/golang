package main

import (
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
)

//ApplicationUsers will represent single user with id, name, balance and version
type ApplicationUsers struct {
	ID      int
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"` // Optimistic Locking
}

var engine *xorm.Engine

func main() {

	engine, err := xorm.NewEngine("postgres", "host=localhost user=postgres password= dbname=postgres sslmode=disable")
	if err != nil {
		log.Panic("Unable to create xorm database engine.")
	}
	engine.Exec("SET SEARCH_PATH TO \"gopay\"")
	engine.Sync(new(ApplicationUsers))
	defer engine.Close()
	fmt.Println("Hello World!")
}
