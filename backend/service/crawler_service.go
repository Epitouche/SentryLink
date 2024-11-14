package service

import (
	"fmt"
	"sync"
)

type UrlsFetched struct {
	mu      sync.Mutex
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
