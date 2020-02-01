package app

import (
	"github.com/lukavavetic/bookstore-users-api/controllers/ping"
	"github.com/lukavavetic/bookstore-users-api/controllers/users"
)

func routes() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
