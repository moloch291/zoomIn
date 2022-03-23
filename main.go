package main

import (
	"zoomIn/magnify"

	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 4 && os.Args[1] == "--magnify" {
		magnify.App()
	} else {
		fmt.Println(
`
Args missing (or too much provided)!

### Usage:

# 1st arg: path to file you want to read
# 2nd arg: string to search for in the file`)
	}
}
