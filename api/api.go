package api

import (
	"fabric-ca-smurf/casdk"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	casdk *casdk.CaSdk
}

func NewApiController() *ApiController {
	sdk := casdk.NewCaSdk()
	return &ApiController{
		casdk: sdk,
	}
}

func (a *ApiController) HandleCaInfo(ctx *gin.Context) {
	cainfo, err := a.casdk.GetCAInfo()
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, cainfo)
}

func (a *ApiController) HandleEnroll(ctx *gin.Context) {
	err := a.casdk.Enroll()
	if err != nil {
		ctx.JSON(http.StatusOK, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}
