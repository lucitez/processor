package process

import (
	"fmt"
	"net/http"
	"time"

	"github.com/lucitez/processor/util"
)

// implements task
type PingNetwork struct {
	url              string
	responseCodes    *util.SafeMapCounter[int]
	urlsVisited      *util.SafeMapCounter[string]
	urlResponseTimes *util.SafeMapCounter[string]
	errCount         *util.SafeMapCounter[string]
}

// next step, record average response time for each url
func (ie *PingNetwork) process() {
	start := time.Now()

	conn, err := http.Get(ie.url)

	if err != nil {
		ie.errCount.Store(err.Error(), 1)

		return
	}
	ie.urlsVisited.Store(ie.url, 1)

	end := time.Now()
	responseTime := end.Sub(start).Milliseconds()

	ie.urlResponseTimes.Store(ie.url, int(responseTime))
	ie.responseCodes.Store(conn.StatusCode, 1)
}

// implements processor
type PingNetworkProcessor struct {
	responseCodes    *util.SafeMapCounter[int]
	urlsVisited      *util.SafeMapCounter[string]
	urlResponseTimes *util.SafeMapCounter[string]
	errors           *util.SafeMapCounter[string]
}

func (ief *PingNetworkProcessor) Setup() {
	ief.responseCodes = &util.SafeMapCounter[int]{SMap: make(map[int]int)}
	ief.errors = &util.SafeMapCounter[string]{SMap: make(map[string]int)}
	ief.urlsVisited = &util.SafeMapCounter[string]{SMap: make(map[string]int)}
	ief.urlResponseTimes = &util.SafeMapCounter[string]{SMap: make(map[string]int)}
}

func (ief PingNetworkProcessor) create(str string) (task, error) {
	return &PingNetwork{
		url:              str,
		responseCodes:    ief.responseCodes,
		errCount:         ief.errors,
		urlResponseTimes: ief.urlResponseTimes,
		urlsVisited:      ief.urlsVisited,
	}, nil
}

func (ief PingNetworkProcessor) print() {
	codes := ief.responseCodes.Get()
	errors := ief.errors.Get()
	urls := ief.urlsVisited.Get()

	fmt.Printf("\nErrors\n\n")

	for err, quantity := range errors {
		fmt.Printf("Found %d instances of error: \"%s\"\n", quantity, err)
	}

	fmt.Printf("\nCodes\n\n")

	for code, quantity := range codes {
		fmt.Printf("Found %d responses with code %d\n", quantity, code)
	}

	fmt.Printf("\nAverage latency\n\n")

	// TODO sort by latency
	for url, timesVisited := range urls {
		averageResponseTime := float32(ief.urlResponseTimes.Access(url)) / float32(timesVisited)

		fmt.Printf("%s: %.2fms\n", url, averageResponseTime)
	}
}
