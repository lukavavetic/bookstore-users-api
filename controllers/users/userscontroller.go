package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(c *gin.Context) {

}


func FindUser() {

}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "sorry")
}