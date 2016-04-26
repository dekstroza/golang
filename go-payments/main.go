package main

import (
	"fmt"
	"net/http"

	"github.com/dekstroza/golang/go-payments/contexts"
	"github.com/dekstroza/golang/go-payments/database"
	"github.com/dekstroza/golang/go-payments/utils"
)

func main() {
	cmdLineFlags := utils.ParseCmdArgs()
	defer database.InitDB(&cmdLineFlags).Close()
	router := contexts.CreateRestServer(cmdLineFlags.BindAddress, cmdLineFlags.BindPort)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *cmdLineFlags.BindAddress, *cmdLineFlags.BindPort), router)

}
