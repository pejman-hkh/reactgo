package main

import (
	"fmt"
	"goreact/app/controller"
	goreact "goreact/src"

	"github.com/gin-gonic/gin"
)

func main() {

	gr := goreact.GoReact{}
	gr.Init()
	r := gin.Default()
	index := &controller.IndexController{}
	fmt.Printf("%d\n", index)
	user := &controller.UserController{}
	r.GET("/", func(ctx *gin.Context) {
		gr.HandleRoute(ctx, index, "Index")
	})

	r.GET("/login", func(ctx *gin.Context) {
		gr.HandleRoute(ctx, user, "Login")
	})

	r.POST("/login", user.LoginPost)

	r.Run(":8090")

}
