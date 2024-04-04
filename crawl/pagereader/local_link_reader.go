package pagereader

import (
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type LocalLinkReader struct {
	parsedUrl *url.URL
	getPage   func(string) (*http.Response, error)
}

func NewLocalLinkReader(rawUrl string, getPage func(string) (*http.Response, error)) (*LocalLinkReader, error) {
	parsedUrl, err := url.Parse(rawUrl)

	if err != nil {
		return nil, err
	}

	return &LocalLinkReader{parsedUrl, getPage}, nil
}

func (pr *LocalLinkReader) Tokenizer() (tokenizer *html.Tokenizer, close func()) {
	resp, err := pr.getPage(pr.parsedUrl.String())

	if err != nil {
		panic(err)
	}

	return html.NewTokenizer(resp.Body), func() {
		resp.Body.Close()
	}
}

func (pr *LocalLinkReader) Read(processUrl func(string)) []string {
	links := []string{}

	getHrefs(pr, func(href string) {
		// if it's prefixed with just a /, it is a local link with same host
		if href != "" && href[0] == '/' {
			newUrl := *pr.parsedUrl

			newUrl.Path = sanitizeLocalHref(href)

			processUrl(newUrl.String())

			links = append(links, newUrl.String())
		}
	})

	return links
}
