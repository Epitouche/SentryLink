package service

// import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// type result struct {
// 	body string
// 	urls []string
// }

// type fetcherService map[string]*result

// func (f fetcherService) Fetch(url string) (string, []string, error) {
// 	if res, ok := f[url]; ok {
// 		return res.body, res.urls, nil
// 	}
// 	return "", nil, fmt.Errorf("not found: %s", url)
// }
