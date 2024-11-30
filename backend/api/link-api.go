package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/controller"
	"github.com/Tom-Mendy/SentryLink/schemas"
)

type LinkApi struct {
	linkController controller.LinkController
}

func NewLinkAPI(
	linkController controller.LinkController,
) *LinkApi {
	return &LinkApi{
		linkController: linkController,
	}
}

func (api *LinkApi) GetLink(ctx *gin.Context) {
	ctx.JSON(200, api.linkController.FindAll())
}

func (api *LinkApi) CreateLink(ctx *gin.Context) {
	err := api.linkController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}

func (api *LinkApi) UpdateLink(ctx *gin.Context) {
	err := api.linkController.Update(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}

func (api *LinkApi) DeleteLink(ctx *gin.Context) {
	err := api.linkController.Delete(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &schemas.Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &schemas.Response{
			Message: "Success!",
		})
	}
}
