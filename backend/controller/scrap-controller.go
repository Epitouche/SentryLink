package controller

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/Tom-Mendy/SentryLink/service"
)

type ScrapController interface {
	Scrap(ctx *gin.Context) []string
}

type scrapController struct {
	service service.ScrapService
}

var validateScrap *validator.Validate

func NewScrapController(scrapService service.ScrapService) ScrapController {
	validateScrap = validator.New()
	return &scrapController{
		service: scrapService,
	}
}


func ExtractLinks(pageURL string) ([]string, error) {
	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the page: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch the page: %s", resp.Status)
	}

	links := []string{}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						// Resolve relative URLs
						parsedURL, err := url.Parse(attr.Val)
						if err != nil {
							continue
						}
						base, err := url.Parse(pageURL)
						if err != nil {
							continue
						}
						links = append(links, base.ResolveReference(parsedURL).String())
					}
				}
			}
		}
	}

	return links, nil
}

func (controller *scrapController) Scrap(ctx *gin.Context) []string {
	pageURL := ctx.Query("url")
	links, err := ExtractLinks(pageURL)
	if err != nil {
		log.Println("Error:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}

	return links
	// ctx.JSON(http.StatusOK, gin.H{"links": links})
}