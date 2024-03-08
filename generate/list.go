package generate

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func List(gen Generator, filename string, numLines int) error {
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

			writeToFile(line, *file)
			wg.Done()
		}(i)
	}

	wg.Wait()

	end := time.Now()

	fmt.Printf("Wrote file in %d milliseconds\n", end.Sub(start).Milliseconds())

	return nil
}
