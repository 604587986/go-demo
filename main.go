package main

import (
	"github.com/gin-gonic/gin"
	"go/api"
	"go/database"
)


func main() {
	database.Connect()

	defer database.Db.Close()


	r := gin.Default()

	r.GET("/users/:id", api.GetUsers)

	r.Run(":8088")


}
