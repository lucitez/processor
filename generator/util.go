package generator

import (
	"os"
)

func WriteToFile(message string, file os.File) {
	file.Write([]byte(message))
}
