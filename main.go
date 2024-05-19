package main

import (
	"github.com/rikiulirniam/RestAPI-Go-Lang/controllers/User/UserController"
	"github.com/rikiulirniam/RestAPI-Go-Lang/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default();
	models.ConnectDatabase()

	r.GET("/api/users", UserController.index)
	r.GET("/api/:id", UserController.show)
	r.POST("/api/users", UserController.store)
	r.PUT("/api/users", UserController.update)
	r.DELETE("/api/users", UserController.destroy)

	r.Run()
}