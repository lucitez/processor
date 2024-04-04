package processor

import (
	"fmt"
	"strconv"

	"github.com/lucitez/processor/util"
)

// implements task
type CountEvens struct {
	value   int
	counter *util.SafeCounter
}

func (ie CountEvens) process() {
	if ie.value%2 == 0 {
		ie.counter.Inc()
	}
}

// implements processor
type CountEvensProcessor struct {
	evenCount *util.SafeCounter
}

func (ief *CountEvensProcessor) Setup() {
	ief.evenCount = &util.SafeCounter{}
}

func (ief CountEvensProcessor) create(str string) (task, error) {
	lineAsInt, err := strconv.Atoi(str)

	if err != nil {
		return nil, err
	}

	return CountEvens{lineAsInt, ief.evenCount}, nil
}

func (ief CountEvensProcessor) print() {
	fmt.Printf("Found %d evens\n", ief.evenCount.Value())
}
