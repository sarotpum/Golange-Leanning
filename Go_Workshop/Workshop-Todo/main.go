package main

import (
	"net/http"

	"github.com/todo/database"

	"github.com/todo/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	database.ConnectDatabase()

	r.GET("/", api.GetAllLists)
	r.GET("/user", api.GetTodoList)
	r.POST("/", api.CreateTodoList)
	r.DELETE("/delete/:id", api.DeleteList)
	r.POST("/upload", api.Upload)
	r.StaticFS("/file", http.Dir("public"))

	r.Run()
}
