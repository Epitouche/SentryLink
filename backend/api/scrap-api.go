package api

import (
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

// func (api *ScrapApi) Scrap(ctx *gin.Context) (string, error) {

// }