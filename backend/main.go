package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	// import the UrlsFetched struct from the service package of my project
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

var db = make(map[string]string)

// -------------------Code that needs to be removed later-------------------
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type result struct {
	body string
	urls []string
}

type fetcherService map[string]*result


func (f fetcherService) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type UrlsFetched struct {
	mu sync.Mutex
	fetched map[string]bool
}

func Crawl(url string, depth int, fetcher Fetcher, urlsFetched *UrlsFetched) {
	if depth <= 0 {
		return
	}

	urlsFetched.mu.Lock()
	if urlsFetched.fetched[url] {
		urlsFetched.mu.Unlock()
		return
	}
	urlsFetched.fetched[url] = true
	urlsFetched.mu.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			Crawl(u, depth-1, fetcher, urlsFetched)
		}(u)
	}
	wg.Wait()
}

// -------------------Code that needs to be removed later-------------------


func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()


	// Crawler test
	r.GET("/crawl/:link", func(c *gin.Context) {

		urls := &UrlsFetched{
			fetched: make(map[string]bool),
		}
		param := c.Params.ByName("link")

		// Temporary code to test the crawler and need to be removed

		var fetcher = fetcherService{
			"https://golang.org/": &result{
				"The Go Programming Language",
				[]string{
					"https://golang.org/pkg/",
					"https://golang.org/cmd/",
				},
			},
			"https://golang.org/pkg/": &result{
				"Packages",
				[]string{
					"https://golang.org/",
					"https://golang.org/cmd/",
					"https://golang.org/pkg/fmt/",
					"https://golang.org/pkg/os/",
				},
			},
			"https://golang.org/pkg/fmt/": &result{
				"Package fmt",
				[]string{
					"https://golang.org/",
					"https://golang.org/pkg/",
				},
			},
			"https://golang.org/pkg/os/": &result{
				"Package os",
				[]string{
					"https://golang.org/",
					"https://golang.org/pkg/",
				},
			},
		}


		Crawl("https://golang.org/", 4, fetcher, urls)
		c.JSON(http.StatusOK, gin.H{"link": param})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// get URL input frontend
	r.POST("/url", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
