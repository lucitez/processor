package generate

import (
	"os"
)

func writeToFile(message string, file os.File) {
	file.Write([]byte(message))
}
