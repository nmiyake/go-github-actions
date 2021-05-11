package main

import (
	"fmt"

	"golang.org/x/tools/imports"
)

func main() {
	_ = imports.Debug
	fmt.Println("Hello, world!")
}
