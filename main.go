package main

import (
	"strconv"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileName, stringToSearchFor := os.Args[1], os.Args[2]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		deliverErrorMessage()
		os.Exit(0)
	}
	defer file.Close()
	output := getOccurrences(file, stringToSearchFor)
	printResults(output, stringToSearchFor, fileName)
}

func deliverErrorMessage() {
	fmt.Println("Unable to search...")
}

func getOccurrences(file *os.File, stringToSearchFor string) string {
	scanner := bufio.NewScanner(file)
	lineNumbers := ""
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), stringToSearchFor) {
			lineNumbers += strconv.Itoa(lineNumber) + ", "
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		deliverErrorMessage()
		os.Exit(1)
	}
	return lineNumbers
}

func printResults(output string, stringToSearchFor string, fileName string) {
	fmt.Println("\nSearch done! Results:")
	if output != "" {
		fmt.Println(
			"\n'" + stringToSearchFor + "' " + "is presented in line(s) " + output[:len(output) - 2] + " in " + fileName)
	} else {
		fmt.Println("The file doesn't contain", "'", stringToSearchFor, "'...")
	}
}
