package main

import (
	"fabric-ca-smurf/api"
	. "fabric-ca-smurf/config"
	. "fabric-ca-smurf/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	apiController := api.NewApiController()
	if err := apiController.Init(); err != nil {
		panic(err)
	}

	MyLogger.Info("ApiController init success.")

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cainfo", apiController.HandleCaInfo)
		v1.POST("/enroll", apiController.HandleEnroll)
	}
	router.Run(fmt.Sprintf(":%d", Configer.GetInt("server.port")))
}
