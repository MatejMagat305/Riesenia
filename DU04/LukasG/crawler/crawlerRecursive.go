package crawler

import "fmt"

func crawlPageR(url string, depth int) *Urls {
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[%d:%s] \n", depth, url)
	}
	return &Urls{depth + 1, urls}
}

func CrawlRecursive(url string, depth int, maxDepth int) {
	if depth <= maxDepth {
		visited[url] = true
		subUrls := crawlPageR(url, depth)
		if depth < maxDepth {
			for _, url := range subUrls.subUrls {
				if _, seen := visited[url]; seen {
					continue
				}
				CrawlRecursive(url, depth + 1, maxDepth)
			}
		}
	}
}
