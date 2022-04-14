package countln

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Directory struct {
	Name       string
	Files      int
	Whitespace int
	Comment    int
	Code       int
	Function   int
}

func (d *Directory) SetName(name string) {
	d.Name = name
}

func (d *Directory) IncrFile() {
	d.Files++
}

func (d *Directory) Update(fileInfo File) {
	d.Whitespace = d.Whitespace + fileInfo.Whitespace
	d.Comment = d.Comment + fileInfo.Comment
	d.Code = d.Code + fileInfo.Code
	d.Function = d.Function + fileInfo.Function
}

func (d *Directory) Print() {
	fmt.Println("-- DIRECTORY COUNTER --")
	fmt.Println("name: ", d.Name)
	fmt.Println("files: ", d.Files)
	fmt.Println("whitespace: ", d.Whitespace)
	fmt.Println("comment: ", d.Comment)
	fmt.Println("code: ", d.Code)
	fmt.Println("function: ", d.Function)
	total := d.Whitespace + d.Comment + d.Code
	fmt.Println("total: ", total)
}

func countDir(directory string, globalCounter *Global) {
	dirCounter := Directory{}
	dirCounter.SetName(directory)

	globalCounter.IncrFolders()

	items, _ := ioutil.ReadDir(directory)
	for _, item := range items {
		if item.IsDir() {
			// handle directory
			dotFolders := strings.Split(item.Name(), ".")
			if dotFolders[0] == "" {
				continue
			}

			target := directory + "/" + item.Name()
			countDir(target, globalCounter)
		} else {
			// handle file
			target := directory + "/" + item.Name()
			fileInfo := countFile(target, globalCounter)
			dirCounter.IncrFile()
			dirCounter.Update(fileInfo)
		}
	}

	// dirCounter.Print()
	// fmt.Println()
}
