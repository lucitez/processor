package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	crawl "github.com/lucitez/processor/crawler"
	"github.com/lucitez/processor/generator"
)

var TEST_URL = "http://go.dev"

func main() {
	before := time.Now()

	urlChan := make(chan string)

	fmt.Println("crawling...")

	go crawl.Crawl(TEST_URL, 0, urlChan, &sync.Map{})

	file, err := os.Create("files/crawl.txt")

	if err != nil {
		panic(err)
	}

	file.Truncate(0) // clear file

	for url := range urlChan {
		generator.WriteToFile(fmt.Sprintf("%s\n", url), *file)
	}

	duration := time.Since(before)

	fmt.Printf("Wrote file in %f seconds", duration.Seconds())
}
