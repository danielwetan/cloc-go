package countln

import (
	"log"
	"os"
)

// Determine is a path are directory or file
func Count(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		// case path is directory
		// dirInfo := countDir(path)
		// dirInfo.Print()
		_ = countDir(path)
		return
	case mode.IsRegular():
		// case path is file
		fileInfo := countFile(path)
		fileInfo.Print()
		return
	}
}
