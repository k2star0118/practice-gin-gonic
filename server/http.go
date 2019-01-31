package server

import (
	"fmt"
	"gin-gonic/config"
	er "gin-gonic/error"
	"gin-gonic/handler"
	"github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

func NewHttpServer(conf *config.AppConfig, waitGroup *sync.WaitGroup) *http.Server {
	engine := gin.Default()
	engine.Use(limit.MaxAllowed(config.Get().MaxConnCount))
	CreateRoutes(conf, engine)

	srv := &http.Server{
		Addr:    conf.HttpEndpoint,
		Handler: engine,
	}

	go func() {
		defer waitGroup.Done()

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Print(fmt.Sprintf("Failed to serve Http request: %+v", err))
			panic(er.HttpServerStart)
		}
	}()

	return srv
}

func CreateRoutes(conf *config.AppConfig, engine *gin.Engine) {
	{
		example := engine.Group("/example")
		{
			example.GET("", handler.GetExample)
		}
	}
}
