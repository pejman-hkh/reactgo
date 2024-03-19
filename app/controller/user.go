package controller

import (
	"goreact/app/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (co *UserController) Register(c *gin.Context) string {
	v := view.UserRenderer{}
	return v.Register()
}

func (co *UserController) Login(c *gin.Context) string {

	v := view.UserRenderer{}
	return v.Login()

}

func (co *UserController) LoginPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "test",
	})
}
