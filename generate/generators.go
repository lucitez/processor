package generate

import (
	"fmt"
	"math/rand"
)

type OddEvenGenerator struct{}

func (oeg OddEvenGenerator) createLine() string {
	return fmt.Sprintf("%d\n", rand.Intn(10))
}

type UrlGenerator struct{}

func (ug UrlGenerator) createLine() string {
	urls := []string{"http://www.google.com/robots.txt", "http://golang.org", "http://dave.com", "http://ziprecruiter.com", "http://bogusurl.org", "http://google.com/invalid-url"}

	return fmt.Sprintf("%s\n", urls[rand.Intn(len(urls))])
}
