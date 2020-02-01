package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lukavavetic/bookstore-users-api/domain"
	"github.com/lukavavetic/bookstore-users-api/services"
	"github.com/lukavavetic/bookstore-users-api/utils/errors"

	//"encoding/json"
	//"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user domain.User

	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	//handle error
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	//handle error
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequest("Invalid JSON body!")

		c.JSON(restErr.Status, restErr)
		return
	}

	res, err := services.CreateUser(user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}


func FindUser() {

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "sorry")
}