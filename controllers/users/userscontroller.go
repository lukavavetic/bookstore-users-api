package users

import (
	"github.com/gin-gonic/gin"
	"github.com/lukavavetic/bookstore-users-api/domain/users"
	"github.com/lukavavetic/bookstore-users-api/services"
	"github.com/lukavavetic/bookstore-users-api/utils/errors"
	"strconv"

	//"encoding/json"
	//"io/ioutil"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User

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

	res, err := services.CreateUser(&user)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusCreated, res)
}


func FindUser() {

}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		err := errors.BadRequest("Param should be type of int64")
		c.JSON(err.Status, err)
		return
	}

	res, getErr := services.GetUser(userId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, res)
}