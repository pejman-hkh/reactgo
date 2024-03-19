package main

import (
	"goreact/app/controller"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Invoke(obj any, name string, args ...any) []reflect.Value {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	return reflect.ValueOf(obj).MethodByName(name).Call(inputs)
}

func RouteHandle(ctx *gin.Context, obj any, method string) {
	rf := Invoke(obj, method, ctx)
	res := rf[0].Interface().(string)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(res))
}

func main() {

	r := gin.Default()
	index := &controller.IndexController{}
	user := &controller.UserController{}
	r.GET("/", func(ctx *gin.Context) {
		RouteHandle(ctx, index, "Index")
	})

	r.GET("/login", func(ctx *gin.Context) {
		RouteHandle(ctx, user, "Login")
	})

	r.POST("/login", func(ctx *gin.Context) {

	})

	r.Run(":8090")

}
