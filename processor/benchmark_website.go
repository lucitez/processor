package processor

import (
	"fmt"
)

// implements task
type UrlBenchmark struct {
	value string
}

func (ub UrlBenchmark) process() {}

// implements processor
type BenchmarkProcessor struct{}

func (bp BenchmarkProcessor) Setup() {}

func (bp BenchmarkProcessor) create(url string) (task, error) {

	return UrlBenchmark{url}, nil
}

func (bp BenchmarkProcessor) print() {
	fmt.Printf("TODO implement me\n")
}
