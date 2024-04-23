package generator

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/lucitez/processor/crawler"
)

type WebLinksGenerator struct {
	BaseUrl string
}

func (wlg WebLinksGenerator) WriteData(filename string) error {

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	file.Truncate(0) // clear file

	start := time.Now()

	urlChan := make(chan string)

	fmt.Println("crawling...")

	go crawler.Crawl(wlg.BaseUrl, 0, urlChan, &sync.Map{})

	for url := range urlChan {
		WriteToFile(fmt.Sprintf("%s\n", url), *file)
	}

	fmt.Printf("Wrote file in %d milliseconds\n", time.Since(start).Milliseconds())

	return nil
}
