package handler

import (
	"gin-gonic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExample(ctx *gin.Context) {
	message := "Successful get message"
	ret, err := service.GetExampleService().GetExample()
	if err != nil {
		ctx.JSON(http.StatusNotFound, message)
		return
	}

	ctx.JSON(http.StatusOK, ret)
}
