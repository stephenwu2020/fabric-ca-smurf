package main

import (
	"fabric-ca-smurf/api"
	. "fabric-ca-smurf/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	apiController := api.NewApiController()

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/cainfo", apiController.HandleCaInfo)
	}
	router.Run(fmt.Sprintf(":%d", Configer.GetInt("server.port")))
}
