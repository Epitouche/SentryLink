package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
)

type ScrapApi struct {
	scrapController controller.ScrapController
}

func NewScrapApi(scrapController controller.ScrapController) ScrapApi {
	return ScrapApi{
		scrapController: scrapController,
	}
}

func (api *ScrapApi) GetScrappedUrl(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, api.scrapController.Scrap(ctx))
}