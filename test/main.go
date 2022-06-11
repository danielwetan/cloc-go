package main

import (
	"github.com/danielwetan/countln"
)

func main() {
	p := "../../countln"
	// p := "./main.go"
	// p := "/home/daniel/code/office/eureka/damascus"

	countln.Count(p)
}
