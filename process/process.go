package process

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type task interface {
	process()
}

type Processor interface {
	Setup()
	create(str string) (task, error)
	print()
}

func Run(p Processor, filename string, concurrency int) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	start := time.Now()

	tasks := make(chan task)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			task, err := p.create(scanner.Text())

			if err != nil {
				fmt.Println(err)
			}

			tasks <- task
		}

		close(tasks)
		wg.Done()
	}()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for t := range tasks {
				t.process()
			}
			wg.Done()
		}()
	}

	wg.Wait()

	end := time.Now()

	fmt.Printf("Processed in %d millis\n", end.Sub(start).Milliseconds())
	p.print()
}
