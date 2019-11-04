package crawler

import "fmt"

func crawlPage(url string, depth int) {
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("[%d:%s] \n", depth, url)
	}
	globalQueueOfUrls <- Urls{depth + 1, urls}
}

func CrawlConcurrent(url string, depth int) {
	totalRuns++
	visited[url] = true
	go crawlPage(url, 0)
	for totalRuns > 0 {
		totalRuns--
		next := <-globalQueueOfUrls
		if next.depth > depth {
			continue
		}
		for _, url := range next.subUrls {
			if _, seen := visited[url]; seen {
				continue
			}
			visited[url] = true
			if next.depth < depth {
				totalRuns++
				go crawlPage(url, next.depth)
			}
		}
	}
}
