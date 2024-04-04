package generator

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LineWriter interface {
	createLine() string
}

func writeLines(gen LineWriter, filename string, numLines int) error {
	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	file.Truncate(0) // clear file

	start := time.Now()

	wg := sync.WaitGroup{}

	for i := 0; i < numLines; i++ {
		wg.Add(1)
		go func(ind int) {
			line := gen.createLine()

			WriteToFile(line, *file)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("Wrote file in %d milliseconds\n", time.Since(start).Milliseconds())

	return nil
}
