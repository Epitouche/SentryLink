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

type controller struct {
	service service.LinkService
}

var validate *validator.Validate

func New(service service.LinkService) LinkController {
	validate = validator.New()
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []schemas.Link {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var link schemas.Link
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		return err
	}

	err = validate.Struct(link)
	if err != nil {
		return err
	}
	err = c.service.Save(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
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

	err = validate.Struct(link)
	if err != nil {
		return err
	}
	c.service.Update(link)
	return nil
}

func (c *controller) Delete(ctx *gin.Context) error {
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

func (c *controller) ShowAll(ctx *gin.Context) {
	links := c.service.FindAll()
	data := gin.H{
		"title": "Link Page",
		"links": links,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
