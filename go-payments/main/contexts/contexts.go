package contexts

import (
	"encoding/json"
	"fmt"

	"github.com/dekstroza/golang/go-payments/main/database"
	"github.com/dekstroza/golang/go-payments/main/models"
	"github.com/gocraft/web"
)

//Context structure containing number of counts a message should be repeated.
type Context struct {
}

//Root just print root
func (c *Context) Root(rw web.ResponseWriter, r *web.Request) {
	fmt.Println("Root")
}

//List location on my router
func (c *Context) List(rw web.ResponseWriter, req *web.Request) {
	rw.Header().Set("Content-Type", "application/json")
	users := []models.ApplicationUser{}
	err := database.DB.Find(&users).Error
	if err != nil {
		fmt.Println(err)
	}
	something, _ := json.Marshal(users)
	fmt.Fprint(rw, string(something))
}
