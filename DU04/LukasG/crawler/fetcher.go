package crawler

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"regexp"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type RealFetcher struct{}

func (f *RealFetcher) Fetch(url string) (body string, urls []string, err error) {
	response, err := http.Head(url)
	if err != nil {
		return url, urls, err
	}
	if response.Status != "200 OK" {
		return url, urls, fmt.Errorf(response.Status)
	}
	response, err = http.Get(url)
	if err != nil {
		return url, urls, err
	}
	_, _ = httputil.DumpResponse(response, false)

	reader := bufio.NewReader(response.Body)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		stringLine := string(line)
		body += stringLine
		var httpRef = regexp.MustCompile(`\s*(?i)href\s*=\s*"[^'"]*`)
		var httpLink = regexp.MustCompile(`https?://(www\.)?[^"']+\.[a-zA-Z0-9()]{1,6}[^?]*`)
		matches := httpRef.FindAllString(stringLine, -1)
		for _, match := range matches {
			match = httpLink.FindString(match)
			if len(match) > 0 {
				urls = append(urls, match)
			}
		}
	}
	return body, urls, nil
}

var realFetcher *RealFetcher = new(RealFetcher)

type NotHttpRefFetcher struct{}

func (f *NotHttpRefFetcher) Fetch(url string) (body string, urls []string, err error) {
	response, err := http.Head(url)
	if err != nil {
		return url, urls, err
	}
	if response.Status != "200 OK" {
		return url, urls, fmt.Errorf(response.Status)
	}
	response, err = http.Get(url)
	if err != nil {
		return url, urls, err
	}
	_, _ = httputil.DumpResponse(response, false)

	reader := bufio.NewReader(response.Body)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		stringLine := string(line)
		body += stringLine
		var httpLink = regexp.MustCompile(`https?://(www\.)?[^"']+\.[a-zA-Z0-9()]{1,6}[^?"']*`)
		matches := httpLink.FindAllString(stringLine, -1)
		for _, match := range matches {
			urls = append(urls, match)
		}
	}
	return body, urls, nil
}

var notRefFetcher *NotHttpRefFetcher = new(NotHttpRefFetcher)

// FakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

type FakeResult struct {
	body string
	urls []string
}

func (f *FakeFetcher) Fetch(url string) (body string, urls []string, err error) {
	res, ok := (*f)[url]
	if ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("URL not found: %s", url)
}

// fetcher is a populated FakeFetcher.
var fakeFetcher = &FakeFetcher{
	"http://golang.org/": &FakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &FakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &FakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &FakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
