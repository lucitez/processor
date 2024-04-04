package generator

import (
	"fmt"
	"math/rand"
)

type UrlGenerator struct {
	NumLines int
}

func (ug UrlGenerator) WriteData(filename string) error {
	return writeLines(ug, filename, ug.NumLines)
}

func (ug UrlGenerator) createLine() string {
	urls := []string{"http://www.google.com/robots.txt", "http://golang.org", "http://dave.com", "http://ziprecruiter.com", "http://bogusurl.org", "http://google.com/invalid-url"}

	return fmt.Sprintf("%s\n", urls[rand.Intn(len(urls))])
}
