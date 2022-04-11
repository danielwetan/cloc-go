package countln

import "fmt"

type Global struct {
	Whitespace int
	Comment    int
	Code       int
	Function   int
}

func (c *Global) Update(fileInfo *File) {
	fmt.Println("UPDATE")
	c.Whitespace = c.Whitespace + fileInfo.Whitespace
	c.Comment = c.Comment + fileInfo.Comment
	c.Code = c.Code + fileInfo.Code
	c.Function = c.Function + fileInfo.Function
}

func (c *Global) Print() {
	fmt.Println("-- GLOBAL COUNTER --")
	fmt.Println("whitespace: ", c.Whitespace)
	fmt.Println("comment: ", c.Comment)
	fmt.Println("code: ", c.Code)
	fmt.Println("function: ", c.Function)
	total := c.Whitespace + c.Comment + c.Code
	fmt.Println("total: ", total)
}
