package main

import (
	"goreact/app/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		index := controller.IndexController{}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(index.Index(ctx)))
	})

	r.GET("/login", func(ctx *gin.Context) {
		user := controller.UserController{}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(user.Login(ctx)))
	})

	r.POST("/login", func(ctx *gin.Context) {

	})

	r.Run(":8090")

}
