package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Tom-Mendy/SentryLink/schemas"
)

type PongApi struct {
}

func NewPongApi() *PongApi {
	return &PongApi{}
}

func (api *PongApi) Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &schemas.Response{
		Message: "Pong",
	})
}
