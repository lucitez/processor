package generator

import (
	"fmt"
	"math/rand"
)

type OddEvenGenerator struct {
	NumLines int
}

func (oeg OddEvenGenerator) WriteData(filename string) error {
	return writeLines(oeg, filename, oeg.NumLines)
}

func (oeg OddEvenGenerator) createLine() string {
	return fmt.Sprintf("%d\n", rand.Intn(10))
}
