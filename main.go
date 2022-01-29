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
		fmt.Println("\nArgs missing or too much!\nUsage: 1st arg = fileName, 2nd arg = stringToSearch")
	}
}
