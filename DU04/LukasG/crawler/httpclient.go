package crawler

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"regexp"
)

func Http() {
	url := "http://golang.org/"
	response, err := http.Head(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if response.Status != "200 OK" {
		fmt.Println(response.Status)
		os.Exit(2)
	}
	response, err = http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	_, _ = httputil.DumpResponse(response, false)
	reader := bufio.NewReader(response.Body)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		stringLine := string(line)
		var httpLink = regexp.MustCompile(`https?://(www\.)?[^"']+\.[a-zA-Z0-9()]{1,6}[^?"']*`)
		matches := httpLink.FindAllString(stringLine, -1)
		for _, match := range matches {
			fmt.Println(match)
		}
	}
}
