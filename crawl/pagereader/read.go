package pagereader

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// only take the path before query strings, anchors, etc.
func sanitizeLocalHref(href string) string {
	var re = regexp.MustCompile(`(/[-_\.\da-zA-Z]*)+`)

	matches := re.FindSubmatch([]byte(href))

	match := string(matches[0])
	match = strings.TrimSuffix(match, "/")
	return match
}

type Tokenizerer interface {
	Tokenizer() (tokenizer *html.Tokenizer, close func())
}

func getHrefs(pr Tokenizerer, processHref func(string)) {
	tokenizer, close := pr.Tokenizer()

	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			close()
			return
		case html.StartTagToken:
			tag, hasAttr := tokenizer.TagName()

			if string(tag) == "a" && hasAttr {
				href := getHref(tokenizer)
				processHref(href)
			}
		}
	}
}

func getHref(tokenizer *html.Tokenizer) string {
	for attr, val, next := tokenizer.TagAttr(); true; attr, val, next = tokenizer.TagAttr() {
		if string(attr) == "href" {
			return string(val)
		}

		if !next {
			break
		}
	}

	return ""
}
