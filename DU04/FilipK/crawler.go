// crawler
package main

//casy zalezia dost od pripojenia
//neviem ale preco mam dost ine cisla ako su v komentaroch
//nemali by sa prehladavat 2krat tie iste stranky
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetcher Fetcher

var v map[string]bool = make(map[string]bool)
var mux *sync.Mutex = new(sync.Mutex)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
//pouzivam waitgroup aby som pockal na dokoncenie prace
//mutex aby som si riadil pristup k mnozine spracovanych
//inak iba volam rekurzivne pre kazdu url crawl s nizsou hlbkou
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth < 0 {
		return
	}
	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s \n", url)
	for _, u := range urls {
		mux.Lock()
		if _, ok := v[u]; ok {
			mux.Unlock()
			continue
		}
		v[u] = true
		mux.Unlock()
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

func main() {
	fetcher = realfetcher // tento ide na web

	//url := "http://dai.fmph.uniba.sk/courses/JAVA"
	//url := "http://dai.fmph.uniba.sk/courses/PARA"
	//url := "http://golang.org/"
	url := "http://dai.fmph.uniba.sk/"
	depth := 2
	//------------------------------------------------
	start := time.Now()

	//tu som si potreboval vyrobit waitgroup inak nic nemenim
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl(url, depth, fetcher, &wg)
	wg.Wait()
	// http://dai.fmph.uniba.sk/courses/JAVA
	// size: 1097
	// 8m44.3811769s

	// http://dai.fmph.uniba.sk/courses/PARA
	// size: 236
	// 1m7.1152031s

	// http://golang.org
	// size: 317
	// 2m10.9955306s

	// http://dai.fmph.uniba.sk/
	// ... ?

	//---------------------------------------------------------------------
	// rekurzivy jemne zoprimalizovany crawler, pamata si navstivene stranky

	//CrawlR(url, 0, depth)

	// http://dai.fmph.uniba.sk/courses/JAVA
	// size: 1097
	// 6m47.2779048s

	// http://dai.fmph.uniba.sk/courses/PARA
	// size: 79
	// 1.6984467s

	// http://golang.org
	// size: 163
	// 1m9.2333664s

	//---------------------------------------------------------------------
	// konkurentny crawler

	//Crawl(url, depth)

	// http://dai.fmph.uniba.sk/courses/JAVA
	// size: 1097
	// 1.5778854s

	// http://dai.fmph.uniba.sk/courses/PARA
	// size: 236
	// 808.6346ms

	// http://golang.org
	// size: 317
	// 1.2192617s

	// http://dai.fmph.uniba.sk/
	// size: 2024
	// 4.7803725s
	//---------------------------------------------------------------------
	//-------------------------
	fmt.Printf("visited: %v\n", v)
	fmt.Printf("size: %d\n", len(v))
	fmt.Println(time.Since(start))

}

type RealFetcher struct{}

var realfetcher *RealFetcher = new(RealFetcher)

//pomocou regexu hladam href ak je chyba tak vratim prazdne suburls
//inak dostanem z regexu malou upravou pekne urls
func (f *RealFetcher) Fetch(url string) (body string, urls []string, err error) {
	req, _ := http.NewRequest("POST", url, nil)
	urls = make([]string, 0)
	req.Header.Add("cache-control", "no-cache")
	res, erro := http.DefaultClient.Do(req)
	if erro != nil {
		return "", []string{}, erro
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	re := regexp.MustCompile(`href *= *"http[^-"]+`)
	h := re.FindAll(b, -1)
	for _, i := range h {
		re = regexp.MustCompile(`http[^"]+`)
		akt := re.Find(i)
		urls = append(urls, string(akt))
	}
	return string(b), urls, nil
}
