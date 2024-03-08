package main

import (
	"flag"
	"fmt"

	"github.com/lucitez/processor/generate"
	"github.com/lucitez/processor/process"
)

func main() {
	var gFlag = flag.Bool("g", false, "boolean: generate a new file for this run")

	var pFlag = flag.String("p", "is_even", "is_even; urls")

	var nFlag = flag.Int("n", 1000, "int: number of lines to generate")

	var cFlag = flag.Int("c", 100, "int: concurrency factor")

	flag.Parse()

	var g generate.Generator
	var p process.Processor
	var filename string

	switch *pFlag {
	case "is_even":
		filename = "files/oddeven.txt"
		g = generate.OddEvenGenerator{}
		p = &process.CountEvensProcessor{}
		p.Setup()
	case "urls":
		filename = "files/urls.txt"
		g = generate.UrlGenerator{}
		p = &process.PingNetworkProcessor{}
		p.Setup()
	default:
		fmt.Printf("Error parsing command line arguments.\n")
		flag.Usage()
		return
	}

	if *gFlag {
		generate.List(g, filename, *nFlag)
	}

	process.Run(p, filename, *cFlag)
}
