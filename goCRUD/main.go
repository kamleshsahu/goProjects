package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var libraryInstance LibraryServiceI
var libraryRepoInstance LibraryRepoI

func Init() {
	libraryRepoInstance = GetLibraryRepo()
	libraryInstance = GetLibraryService(libraryRepoInstance)
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	Init()
	router.POST("/book", CreateUser)
	router.POST("/book/:id", UpdateUser)
	router.DELETE("/book/:id", DeleteUser)
	router.GET("/book", GetAllUsers)

	router.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
