package main

import (
	"flag"
	"net/http"
)
import (
	"github.com/dekstroza/golang/go-payments/contexts"
	"github.com/dekstroza/golang/go-payments/database"
)
import (
	"github.com/gocraft/web"
)

func main() {
	dbHostPtr := flag.String("dbHost", "localhost", "Address or name of the database host.")
	dbHostPortPtr := flag.Int("dbPort", 5432, "Port on which database server will accept connections.")
	dbUsernamePtr := flag.String("user", "postgres", "Username for database connection")
	dbPasswordPtr := flag.String("password", "", "Password for database connection")
	dbNamePtr := flag.String("db", "postgres", "Name of the database to connect to.")
	dbSchemaPtr := flag.String("schema", "public", "Name of database schema that will be used.")

	dbMaxIdleConnectionsPtr := flag.Int("maxIdle", 20, "Maximum number of idle connections in the pool.")
	dbMaxConnectionsPtr := flag.Int("maxConnections", 60, "Maximum number of database connections.")

	flag.Parse()
	database.InitDB(dbHostPtr, dbHostPortPtr, dbUsernamePtr, dbPasswordPtr, dbNamePtr, dbSchemaPtr, dbMaxIdleConnectionsPtr, dbMaxConnectionsPtr)

	router := web.New(contexts.Context{})
	router.Get("/", (*contexts.Context).Root)
	router.Get("/users", (*contexts.Context).List)
	router.Get("/users/:id", (*contexts.Context).FindUser)
	router.Post("/users", (*contexts.Context).InsertUser)
	http.ListenAndServe("localhost:3000", router)
}
