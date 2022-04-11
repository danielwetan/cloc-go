package target

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	// file, err := os.Open("/path/to/file.txt")
	// file, err := os.Open("./example_http.go")
	// file, err := os.Open("./example_massive.go")
	// file, err := os.Open("./example.go")

	filename := "./example.go"
	// filename := "./example_http.go"
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var (
		whitespace int
		comment    int
		code       int
		function   int
	)

	nestedComment := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Fields(scanner.Text())

		if len(s) == 0 {
			whitespace++
			continue
		}

		if s[0] == "/*" {
			nestedComment = 1
			continue
		}

		if s[0] == "*/" {
			nestedComment = 2
		}

		if nestedComment == 1 {
			continue
		}

		if nestedComment == 2 {
			comment++
			nestedComment = 0
			continue
		}

		if s[0] == "//" {
			comment++
			continue
		}

		if s[0] == "func" {
			function++
			code++
			continue
		}

		code++
	}

	duration := time.Since(start)
	fmt.Println("whitespace: ", whitespace)
	fmt.Println("comment: ", comment)
	fmt.Println("code: ", code)
	fmt.Println("function: ", function)
	fmt.Println("---------")
	fmt.Println("execution time: ", duration)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
