package crawler

import (
	"fmt"
	"time"
)

var (
	globalQueueOfUrls = make(chan Urls)
	totalRuns         = 0
	visited           = make(map[string]bool)
)

var fetcher Fetcher
func Crawl() {
	//fetcher = fakeFetcher // tento ide do vlastnej pidi datovej struktury
	//fetcher = realFetcher // tento ide na web
	fetcher = notRefFetcher // tento ide na web a nehlada iba linky zacinajuce sa href, ale vsetko co vyzera ako link

	//url := "http://dai.fmph.uniba.sk/courses/JAVA"
	//url := "http://dai.fmph.uniba.sk/courses/PARA"
	url := "http://golang.org/"
	//url := "http://dai.fmph.uniba.sk/"
	depth := 2
	start := time.Now()

	//---------------------------------------------------------------------
	// navštivi viackrat stranky

	//CrawlBasic(url, depth)

	//---------------------------------------------------------------------
	// rekurzivy jemne zoprimalizovany crawler, pamata si navstivene stranky

	//CrawlRecursive(url, 0, depth)

	//---------------------------------------------------------------------
	// konkurentny crawler

	CrawlConcurrent(url, depth)

	//---------------------------------------------------------------------

	fmt.Printf("visited: %v\n", visited)
	fmt.Printf("size: %d\n", len(visited))
	fmt.Println(time.Since(start))
}


