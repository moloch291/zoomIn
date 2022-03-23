package magnify

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func App() {
	fileName, stringToSearchFor := os.Args[2], os.Args[3]
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
	fmt.Println("Unable to search... :(")
}

func getOccurrences(file *os.File, stringToSearchFor string) string {
	scanner := bufio.NewScanner(file)
	lineNumbers := ""
	lineNumber := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), stringToSearchFor) {
			lineNumbers += strconv.Itoa(lineNumber) + ": " + scanner.Text() + "\n"
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		deliverErrorMessage()
		os.Exit(0)
	}
	return lineNumbers
}

func printResults(output string, stringToSearchFor string, fileName string) {
	fmt.Println("'" + stringToSearchFor + "' " + "is presented in " + fileName + " in line(s):\n")
	if output != "" {
		fmt.Println(output[:len(output) - 2])
		fmt.Println("\nSearch done!")
	} else {
		fmt.Println("The file doesn't contain", "'", stringToSearchFor, "'...")
	}
}

