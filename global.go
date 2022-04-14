package countln

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type Global struct {
	Files      int
	Folders    int
	Whitespace int
	Comment    int
	Code       int
	Function   int
}

func (c *Global) Update(fileInfo *File) {
	c.Whitespace = c.Whitespace + fileInfo.Whitespace
	c.Comment = c.Comment + fileInfo.Comment
	c.Code = c.Code + fileInfo.Code
	c.Function = c.Function + fileInfo.Function
	c.Files = c.Files + 1
}

func (c *Global) IncrFolders() {
	c.Folders++
}

func (c *Global) GetFolders() int {
	return c.Folders - 1
}

// func (c *Global) Print() {
// 	fmt.Println("-- GLOBAL COUNTER --")
// 	fmt.Println("files: ", c.Files)
// 	fmt.Println("folders: ", c.GetFolders())
// 	fmt.Println("whitespace: ", c.Whitespace)
// 	fmt.Println("comment: ", c.Comment)
// 	fmt.Println("code: ", c.Code)
// 	fmt.Println("function: ", c.Function)
// 	total := c.Whitespace + c.Comment + c.Code
// 	fmt.Println("total: ", total)
// }

func (g *Global) Print() {
	data := [][]string{
		{strconv.Itoa(g.Files), strconv.Itoa(g.Folders), strconv.Itoa(g.Whitespace), strconv.Itoa(g.Comment), strconv.Itoa(g.Function), strconv.Itoa(g.Code)},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"file", "folder", "whitespace", "comment", "function", "code"})
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.AppendBulk(data)

	fmt.Println()
	table.Render()
}
