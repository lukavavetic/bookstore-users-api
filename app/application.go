package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lukavavetic/bookstore-users-api/datasource/mysql/users_db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	users_db.Init()
	routes()

	router.Run(":8888")
}