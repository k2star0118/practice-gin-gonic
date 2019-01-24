package error

import (
	"errors"
	"gin-gonic/model"
	"github.com/gin-gonic/gin"
)

var (
	HttpServerStart = errors.New("http server start error")
)

func NewHttpError(ctx *gin.Context, status int, err error) {
	if err == nil {
		err = errors.New("")
	}

	response := model.HttpResponse{
		Code:    status,
		Message: err.Error(),
	}

	ctx.JSON(status, response)
}
