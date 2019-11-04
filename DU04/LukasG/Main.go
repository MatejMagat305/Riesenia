package main

import (
	"./crawler"
	"./determinant"
)

func main() {
	crawler.HackImport()
	determinant.HackImport()

	//determinant.Test() //U1
	crawler.Crawl() //U2
}
