package api

import (
	"fabric-ca-smurf/caclient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCaInfo(ctx *gin.Context) {
	cainfo, err := caclient.GetCAInfo()
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, cainfo)
}
