package main

// import (
// 	"fmt"
// 	"sync"
//
// )

// type Fetcher interface {
// 	// Fetch returns the body of URL and
// 	// a slice of URLs found on that page.
// 	Fetch(url string) (body string, urls []string, err error)
// }

// type UrlsFetched struct {
// 	mu sync.Mutex
// 	fetched map[string]bool
// }

// // Crawl uses fetcher to recursively crawl
// // pages starting with url, to a maximum of depth.
// func Crawl(url string, depth int, fetcher Fetcher, urlsFetched *UrlsFetched) {
// 	if depth <= 0 {
// 		return
// 	}

// 	urlsFetched.mu.Lock()
// 	if urlsFetched.fetched[url] {
// 		urlsFetched.mu.Unlock()
// 		return
// 	}
// 	urlsFetched.fetched[url] = true
// 	urlsFetched.mu.Unlock()

// 	body, urls, err := fetcher.Fetch(url)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("found: %s %q\n", url, body)

// 	var wg sync.WaitGroup
// 	for _, u := range urls {
// 		wg.Add(1)
// 		go func(u string) {
// 			defer wg.Done()
// 			Crawl(u, depth-1, fetcher, urlsFetched)
// 		}(u)
// 	}
// 	wg.Wait()
// }

// // func main() {
// // 	urls := &UrlsFetched{fetched:make(map[string]bool)}
// // 	Crawl("https://golang.org/", 4, fetcher, urls)
// // }

// // fakeFetcher is Fetcher that returns canned results.
// type fakeFetcher map[string]*fakeResult

// type fakeResult struct {
// 	body string
// 	urls []string
// }

// func (f fakeFetcher) Fetch(url string) (string, []string, error) {
// 	if res, ok := f[url]; ok {
// 		return res.body, res.urls, nil
// 	}
// 	return "", nil, fmt.Errorf("not found: %s", url)
// }

// // fetcher is a populated fakeFetcher.
// var fetcher = fakeFetcher{
// 	"https://golang.org/": &fakeResult{
// 		"The Go Programming Language",
// 		[]string{
// 			"https://golang.org/pkg/",
// 			"https://golang.org/cmd/",
// 		},
// 	},
// 	"https://golang.org/pkg/": &fakeResult{
// 		"Packages",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/cmd/",
// 			"https://golang.org/pkg/fmt/",
// 			"https://golang.org/pkg/os/",
// 		},
// 	},
// 	"https://golang.org/pkg/fmt/": &fakeResult{
// 		"Package fmt",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// 	"https://golang.org/pkg/os/": &fakeResult{
// 		"Package os",
// 		[]string{
// 			"https://golang.org/",
// 			"https://golang.org/pkg/",
// 		},
// 	},
// }
