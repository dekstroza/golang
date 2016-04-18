package main

import (
	"net/http"

	"github.com/dekstroza/golang/go-payments/main/contexts"
	"github.com/dekstroza/golang/go-payments/main/database"
	"github.com/gocraft/web"
)

func main() {
	database.InitDB()
	database.CreateUser()

	router := web.New(contexts.Context{})
	router.Middleware(web.LoggerMiddleware)
	router.Middleware(web.ShowErrorsMiddleware)
	router.Get("/", (*contexts.Context).Root)
	router.Get("/list", (*contexts.Context).List)
	http.ListenAndServe("localhost:3000", router)
}
