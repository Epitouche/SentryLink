package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
)

type LinkController interface {
	FindAll() []schemas.Link
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type linkController struct {
	service service.LinkService
}

var validateLink *validator.Validate

func NewLinkController(service service.LinkService) LinkController {
	validateLink = validator.New()
	return &linkController{
		service: service,
	}
}

func (c *linkController) FindAll() []schemas.Link {
	return c.service.FindAll()
}

func (c *linkController) Save(ctx *gin.Context) error {
	var link schemas.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		return err
	}

	err = validateLink.Struct(link)
	if err != nil {
		return err
	}
	err = c.service.Save(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *linkController) Update(ctx *gin.Context) error {
	var link schemas.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	link.Id = id

	err = validateLink.Struct(link)
	if err != nil {
		return err
	}

	err = c.service.Update(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *linkController) Delete(ctx *gin.Context) error {
	var link schemas.Link
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	link.Id = id
	err = c.service.Delete(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *linkController) ShowAll(ctx *gin.Context) {
	links := c.service.FindAll()
	data := gin.H{
		"title": "Link Page",
		"links": links,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
