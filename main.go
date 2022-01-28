package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName := os.Args[1]
	stringToSearchFor := os.Args[2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	output := reader(file, stringToSearchFor)
	printResults(output, stringToSearchFor, fileName)
}

func reader(file *os.File, stringToSearchFor string) string {
	output := ""
	scanner := bufio.NewScanner(file)
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), stringToSearchFor) {
			output += string(rune(lineNumber)) + ", "
		}
		lineNumber++

	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return output
}

func printResults(output string, stringToSearchFor string, fileName string) {
	fmt.Println("\nSearch done! Results:")
	if output != "" {
		result := "'" + stringToSearchFor + "'" + "is presented in line(s) " + output[:len(output)-2] + " in" + fileName
		fmt.Println(result)
	} else {
		fmt.Println("The file doesn't contain", "'", stringToSearchFor, "'...")
	}
}
