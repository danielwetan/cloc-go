package countln

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type File struct {
	Name       string
	Language   string
	Whitespace int
	Comment    int
	Code       int
	Function   int
}

func (f *File) Print() {
	fmt.Println("name: ", f.Name)
	fmt.Println("language: ", f.Language)
	fmt.Println("whitespace: ", f.Whitespace)
	fmt.Println("comment: ", f.Comment)
	fmt.Println("code: ", f.Code)
	fmt.Println("function: ", f.Function)
	total := f.Whitespace + f.Comment + f.Code
	fmt.Println("total: ", total)
	fmt.Println()
}

func countFile(filename string, globalCounter Global) File {
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
			comment++
			nestedComment = 1
			continue
		}

		if s[0] == "*/" {
			nestedComment = 2
		}

		if nestedComment == 1 {
			comment++
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

	language := strings.Split(filename, ".")

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return File{}
	}

	f := File{
		Name:       filename,
		Language:   language[len(language)-1],
		Whitespace: whitespace,
		Comment:    comment,
		Code:       code,
		Function:   function,
	}
	globalCounter.Update(&f)

	// f.Print()
	// fmt.Println()

	return f
}
