package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Tom-Mendy/SentryLink/schemas"
	"github.com/Tom-Mendy/SentryLink/service"
)

type GithubTokenController interface {
	FindAll() []schemas.GithubToken
	Save(ctx *gin.Context) error
	Update(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
}

type githubTokenController struct {
	service service.GithubTokenService
}

var validateGithubToken *validator.Validate

func NewGithubTokenController(service service.GithubTokenService) GithubTokenController {
	validateGithubToken = validator.New()
	return &githubTokenController{
		service: service,
	}
}

func (c *githubTokenController) FindAll() []schemas.GithubToken {
	return c.service.FindAll()
}

func (c *githubTokenController) Save(ctx *gin.Context) error {
	var link schemas.GithubToken
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		return err
	}

	err = validateGithubToken.Struct(link)
	if err != nil {
		return err
	}
	err = c.service.Save(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *githubTokenController) Update(ctx *gin.Context) error {
	var link schemas.GithubToken
	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	link.Id = id

	err = validateGithubToken.Struct(link)
	if err != nil {
		return err
	}

	err = c.service.Update(link)
	if err != nil {
		return err
	}
	return nil
}

func (c *githubTokenController) Delete(ctx *gin.Context) error {
	var token schemas.GithubToken
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	token.Id = id
	err = c.service.Delete(token)
	if err != nil {
		return err
	}
	return nil
}

func (c *githubTokenController) ShowAll(ctx *gin.Context) {
	links := c.service.FindAll()
	data := gin.H{
		"title": "Link Page",
		"links": links,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
