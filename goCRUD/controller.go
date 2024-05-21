package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(ctx *gin.Context) {
	var user User
	if err := ctx.BindJSON(&user); err != nil {
		ctx.Error(err)
		return
	}

	fmt.Println(user)

	createdUser, err := libraryInstance.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": createdUser})
}

func UpdateUser(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.Error(err)
		return
	}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(err)
		return
	}

	fmt.Println(user)

	updatedUser, err := libraryInstance.UpdateUser(user.Id, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": updatedUser})
}

func DeleteUser(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindUri(&user); err != nil {
		ctx.Error(err)
		return
	}

	err := libraryInstance.DeleteUser(user.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true})
}

func GetAllUsers(ctx *gin.Context) {

	var params interface{}
	allUsers, err := libraryInstance.GetUsers(params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "data": err})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": allUsers})
}
