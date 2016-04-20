package contexts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dekstroza/golang/go-payments/database"
	"github.com/dekstroza/golang/go-payments/models"
	"github.com/gocraft/web"
	"github.com/twinj/uuid"
)

//Context structure containing number of counts a message should be repeated.
type Context struct {
}

//Root just print root
func (c *Context) Root(rw web.ResponseWriter, r *web.Request) {
	fmt.Println("Root")
}

//FindUser will find user with path param id and return json representation
func (c *Context) FindUser(rw web.ResponseWriter, req *web.Request) {

	userID := req.PathParams["id"]
	ID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		user := models.ApplicationUser{}
		database.DB.First(&user, ID)
		if user.ID == "" {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			jsonUser, _ := json.Marshal(user)
			rw.Header().Set("Content-Type", "application/json")
			fmt.Fprint(rw, string(jsonUser))
		}
	}
}

//List location on my router
func (c *Context) List(rw web.ResponseWriter, req *web.Request) {
	users := []models.ApplicationUser{}
	err := database.DB.Find(&users).Error
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		if len(users) == 0 {
			rw.WriteHeader(http.StatusNotFound)
		} else {
			jsonArray, _ := json.Marshal(users)
			rw.Header().Set("Content-Type", "application/json")
			fmt.Fprint(rw, string(jsonArray))
		}
	}

}

//InsertUser will insert new user from submited json
func (c *Context) InsertUser(rw web.ResponseWriter, req *web.Request) {
	user := models.ApplicationUser{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		user.ID = uuid.NewV4().String()
		database.DB.Create(&user)
		rw.Header().Set("X-Xss-Protection", "1")
		rw.Header().Set("X-Frame-Options", "SAMEORIGIN")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Vary", "Origin; X-Origin")
		rw.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Authorization, Content-Type")
		rw.Header().Set("Connection", "keep-alive")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		rw.Header().Set("Access-Control-Max-Age", "3600")
		rw.Header().Set("Location", "http://"+req.Host+":3000"+req.RequestURI+"/"+user.ID)
		rw.WriteHeader(http.StatusCreated)
	}
}
