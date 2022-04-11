package countln

import (
	"log"
	"os"
)

func Count(path string) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Println(err)
		return
	}

	globalCounter := Global{}

	// Determine is a path are directory or file
	switch mode := fi.Mode(); {
	case mode.IsDir():
		countDir(path, &globalCounter)
	case mode.IsRegular():
		_ = countFile(path, &globalCounter)
	}

	globalCounter.Print()
}
