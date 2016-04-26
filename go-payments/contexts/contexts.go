package contexts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dekstroza/golang/go-payments/database"
	"github.com/dekstroza/golang/go-payments/models"
	"github.com/gocraft/web"
	"github.com/twinj/uuid"
)

//Context structure
type Context struct {
}

var serverPort int
var serverIP string

//CreateRestServer will create gocraft rest router
func CreateRestServer(bindAddress *string, bindPort *int) *web.Router {
	serverPort = *bindPort
	serverIP = *bindAddress
	router := web.New(Context{})
	router.Middleware(web.LoggerMiddleware)
	router.Get("/users", (*Context).List)
	router.Get("/users/:id", (*Context).FindUser)
	router.Post("/users", (*Context).InsertUser)
	return router
}

//FindUser will find user with path param id and return json representation
func (c *Context) FindUser(rw web.ResponseWriter, req *web.Request) {
	userID := req.PathParams["id"]
	user := models.ApplicationUser{}
	database.DB.Where("id = ?", userID).Find(&user)
	if user.ID == "" {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		jsonUser, _ := json.Marshal(user)
		rw.Header().Set("Content-Type", "application/json")
		fmt.Fprint(rw, string(jsonUser))

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

//InsertUser into database using body formated as json
func (c *Context) InsertUser(rw web.ResponseWriter, req *web.Request) {
	user := models.ApplicationUser{}
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		user.ID = uuid.NewV4().String()
		database.DB.Create(&user)
		rw.Header().Set("Location", fmt.Sprintf("http://%s:%d%s%s", serverIP, serverPort, req.RequestURI, user.ID))
		rw.WriteHeader(http.StatusCreated)
	}
}
