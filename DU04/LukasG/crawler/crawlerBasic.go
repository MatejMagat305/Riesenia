package crawler

import "fmt"

// CrawlBasic uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlBasic(url string, depth int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth < 0 {
		return
	}
	visited[url] = true
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("[%d:%s] \n", depth, url)
	if depth > 0 {
		for _, u := range urls {
			CrawlBasic(u, depth - 1)
		}
	}
	return
}
