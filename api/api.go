package api

import (
	"fabric-ca-smurf/caclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	caClient *caclient.CaClient
}

func NewApiController() *ApiController {
	client := caclient.NewCaClient()
	return &ApiController{
		caClient: client,
	}
}

func (a *ApiController) HandleCaInfo(ctx *gin.Context) {
	cainfo, err := a.caClient.GetCAInfo()
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, cainfo)
}
