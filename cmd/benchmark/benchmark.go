package main

import (
	"github.com/lucitez/processor/benchmarker"
)

const URL = "https://go.dev"

func main() {
	b := benchmarker.New(URL)

	b.BenchmarkWebsite()
}
