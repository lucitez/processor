package main

import (
	"fmt"
	"time"

	"github.com/lucitez/processor/crawl"
	"github.com/lucitez/processor/util"
)

var TEST_URL = "https://go.dev"

func main() {
	before := time.Now()

	urlChan := make(chan string)
	safeMap := util.SafeMapCounter[string]{SMap: make(map[string]int)}

	fmt.Println("crawling...")

	go crawl.Crawl(TEST_URL, 0, urlChan, &safeMap)

	urls := []string{}

	for url := range urlChan {
		urls = append(urls, url)
	}

	duration := time.Since(before)

	fmt.Printf("Found %d urls in %f seconds", len(urls), duration.Seconds())
}
