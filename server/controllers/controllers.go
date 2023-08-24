package controllers

import (
	"fmt"
	"net/http"
	dataEncryption "server/utils"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	user := ctx.PostForm("user")
	password := ctx.PostForm("password")
	bytes, _ := dataEncryption.Encryption(password)
	a := dataEncryption.Decrypt(string(bytes), user)

	fmt.Printf("bytes: %v\n", string(bytes))
	fmt.Printf("a: %v\n", a)

	fmt.Printf("user: %v\n", user)
	fmt.Printf("password: %v\n", string(password))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
