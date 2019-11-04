// https://code.google.com/p/go-tour/source/browse/solutions/webcrawler.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type RealFetcher struct{
}
var realfetcher *RealFetcher = new(RealFetcher)

// moj kod
func (f *RealFetcher) Fetch(url1 string) (body string, urls []string, err error) {
	body0 := ""
	urls0:=make([]string,0)
	response, err := http.Head(url1)
	if err!=nil {
		return body0, urls0, errors.New("nenacitalo!")
	}
	for k, v := range response.Header { // Content-Type:
		fmt.Println(k+":", v) } // [text/html; charset=ISO-8859-2]
	response, err = http.Get(url1)
	if err != nil{

		return body0, urls0, errors.New("nenacitalo!")
	}
	reader := bufio.NewReader(response.Body) // čítame telo
	for  {
		line, _, err := reader.ReadLine()
		if err == io.EOF { break }
		strline := string(line)
		body0=body0+strline
		var httpRef = regexp.MustCompile(`\s*(?i)href\s*=\s*(\"([^"]*\")|'[^']*'|([^'">\s]+))`)
		matches := httpRef.FindAllString(strline, -1)
		for _, match := range matches {
			_ , err := url.Parse(match)
			if err != nil {
				urls0 = append(urls0, match)
			}
		}
	}
	return body0, urls0, nil
}
//koniec mojho kodu

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var loading = errors.New("url load in progress")

func Crawl(url string, depth int, realfetcher *RealFetcher, group *sync.WaitGroup) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}
	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		fmt.Printf("<- Done with %v, already fetched.\n", url)
		return
	}
	fetched.m[url] = loading
	fetched.Unlock()
	body, urls, err := realfetcher.Fetch(url)

	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()

	if err != nil {
		fmt.Printf("<- Error on %v: %v\n", url, err)
		return
	}
	fmt.Printf("Found: %s %q\n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		group.Add(1)
		go func(url string) {
			Crawl(url, depth-1, realfetcher, nil)
			done <- true
			group.Done()
		}(u)
	}
		for i := range urls {
			fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls))
			<-done
	}
	fmt.Printf("<- Done with %v\n", url)

}
/*
func main() {
	var wg sync.WaitGroup
	p:=time.Now()
	Crawl("http://dai.fmph.uniba.sk/courses/JAVA", 4, realfetcher,&wg)
	wg.Wait()
	fmt.Println("Fetching stats\n--------------")
	fmt.Println(time.Now().Sub(p),len(fetched.m))
	//for url, err := range fetched.m {
	//	if err != nil {
	//		fmt.Printf("%v failed: %v\n", url, err)
	//	} else {
	//		fmt.Printf("%v was fetched\n", url)
	//	}
	//}

}*/
