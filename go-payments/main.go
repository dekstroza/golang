package main

import (
	"net/http"

	"github.com/dekstroza/golang/go-payments/contexts"
	"github.com/dekstroza/golang/go-payments/database"
	"github.com/gocraft/web"
)

func main() {
	database.InitDB()

	router := web.New(contexts.Context{})
	router.Middleware(web.LoggerMiddleware)
	router.Middleware(web.ShowErrorsMiddleware)
	router.Get("/", (*contexts.Context).Root)
	router.Get("/users", (*contexts.Context).List)
	router.Get("/users/:id", (*contexts.Context).FindUser)
	router.Post("/users", (*contexts.Context).InsertUser)
	http.ListenAndServe("localhost:3000", router)
}
