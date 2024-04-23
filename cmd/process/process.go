package main

import (
	"flag"
	"fmt"

	"github.com/lucitez/processor/generator"
	"github.com/lucitez/processor/processor"
)

func main() {
	var gFlag = flag.Bool("g", false, "boolean: generate a new file for this run")

	var pFlag = flag.String("p", "is_even", "is_even; urls; weblinks")

	var uFlag = flag.String("u", "https://go.dev", "base url to crawl")

	var nFlag = flag.Int("n", 1000, "int: number of lines to generate")

	var cFlag = flag.Int("c", 100, "int: concurrency factor")

	flag.Parse()

	var g generator.Generator
	var p processor.Processor
	var filename string

	switch *pFlag {
	case "is_even":
		filename = "files/oddeven.txt"
		g = generator.OddEvenGenerator{NumLines: *nFlag}
		p = &processor.CountEvensProcessor{}
		p.Setup()
	case "urls":
		filename = "files/urls.txt"
		g = generator.UrlGenerator{NumLines: *nFlag}
		p = &processor.PingNetworkProcessor{}
		p.Setup()
	case "weblinks":
		filename = "files/weblinks.txt"
		g = generator.WebLinksGenerator{BaseUrl: *uFlag}
		p = &processor.BenchmarkProcessor{}
		p.Setup()
	default:
		fmt.Printf("Error parsing command line arguments.\n")
		flag.Usage()
		return
	}

	if *gFlag {
		if err := g.WriteData(filename); err != nil {
			fmt.Printf("Error writing to file, %s\n", err)
		}

	}

	processor.Run(p, filename, *cFlag)
}
