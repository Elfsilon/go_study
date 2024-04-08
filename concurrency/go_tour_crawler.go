package concurrency

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type WebCrawler struct {
	wg    *sync.WaitGroup
	m     *sync.Mutex
	found map[string]struct{}
}

func NewWebCrawler(wg *sync.WaitGroup, m *sync.Mutex) WebCrawler {
	return WebCrawler{
		found: make(map[string]struct{}, 0),
		wg:    wg,
		m:     m,
	}
}

// DONE: Fetch URLs in parallel.
// TODO: Don't fetch the same URL twice.
func (c *WebCrawler) Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	c.m.Lock()
	c.found[url] = struct{}{}
	c.m.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		u := u

		c.m.Lock()
		_, ok := c.found[u]
		c.m.Unlock()

		if ok {
			continue
		}

		c.wg.Add(1)

		go func() {
			defer c.wg.Done()

			c.Crawl(u, depth-1, fetcher)
		}()
	}
}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// The Go Programming Language
// Packages
// cmd: not found
// Package fmt
// Package os
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func RunCrawler() {
	wg := sync.WaitGroup{}
	m := sync.Mutex{}

	c := NewWebCrawler(&wg, &m)

	c.Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}
