package handler

import (
	"gin-gonic/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

func GetExample(ctx *gin.Context) {
	errMessage := "Not implement get example"
	ret, err := service.GetExampleService().GetExample()
	if err != nil {
		ctx.JSON(http.StatusNotFound, errMessage)
		return
	}

	enablePut, ok := os.LookupEnv("ENABLE_GET")
	if ok {
		if strings.ToLower(enablePut) == "true" {
			ctx.JSON(http.StatusOK, ret)
			return
		}
	}

	ctx.JSON(http.StatusForbidden, "Does not enable get method")
}

func PutExample(ctx *gin.Context) {
	message := "Not implement put example"
	ret, err := service.GetExampleService().PutExample()
	if err != nil {
		ctx.JSON(http.StatusNotFound, message)
		return
	}

	enablePut, ok := os.LookupEnv("ENABLE_PUT")
	if ok {
		if strings.ToLower(enablePut) == "true" {
			ctx.JSON(http.StatusOK, ret)
			return
		}
	}

	ctx.JSON(http.StatusForbidden, "Does not enable put method")

}
