package main

import (
	"books-project/app"
	"books-project/auth"
	"books-project/db"
	"books-project/middleware"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine
func main() {
	conn := db.InitDB()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	handler := app.New(conn)

	// endpoint home
	router.GET("/", auth.HomeHandler)

	// login
	router.GET("/login", auth.LoginGetHandler)
	router.POST("/login", auth.LoginPostHandler)

	// Endpoint get all books
	router.GET("/books", middleware.AuthValid, handler.GetBooks) 
	router.Run()
}