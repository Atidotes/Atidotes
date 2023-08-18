package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	user     string
	password string
}

func Login(ctx *gin.Context) {
	user := ctx.PostForm("user")
	password := ctx.PostForm("password")
	fmt.Printf("user: %v\n", user)
	fmt.Printf("password: %v\n", password)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
