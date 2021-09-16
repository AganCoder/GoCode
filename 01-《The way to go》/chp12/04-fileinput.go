package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {

	absPath, _ := filepath.Abs("./README.md")
	fmt.Printf("%v", absPath)

	inputFile, err := os.Open(absPath)

	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		fmt.Printf("%v", err)
		return // exit the function on error
	}
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, err := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)

		if err == io.EOF {
			break
		}
	}
}
