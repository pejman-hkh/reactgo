package controller

import (
	"goreact/app/view"

	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func (co *IndexController) Index(ctx *gin.Context) string {
	v := view.IndexRenderer{}
	return v.Index()
}
