package crawl

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"

	"github.com/lucitez/processor/crawl/pagereader"
	"github.com/lucitez/processor/util"
)

var MAX_DEPTH = 3

func Crawl(url string, depth int, out chan<- string, safemap *util.SafeMapCounter[string]) {
	safemap.Store(url, 1)

	llr, err := pagereader.NewLocalLinkReader(url, http.Get)

	if err != nil {
		fmt.Printf("Unable to create page reader, %s\n", err)
		return
	}

	out <- url

	// Once we reach our max depth, don't bother visiting the url to get child urls
	if depth >= MAX_DEPTH {
		return
	}

	wg := sync.WaitGroup{}
	llr.Read(func(foundUrl string) {
		// do not visit urls for files. TODO move this to sanitize in the reader
		var re = regexp.MustCompile(`.*\.(zip|gz|msi|amd|pkg)`)

		matches := re.Find([]byte(foundUrl))

		if matches != nil || safemap.Access(foundUrl) > 0 {
			return
		}

		wg.Add(1)

		go func() {
			Crawl(foundUrl, depth+1, out, safemap)
			wg.Done()
		}()
	})

	wg.Wait()

	if depth == 0 {
		close(out)
	}
}
