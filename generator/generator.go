package generator

type Generator interface {
	WriteData(filename string) error
}
