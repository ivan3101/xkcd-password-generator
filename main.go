package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := getListOfWords()

	fmt.Println(words)
}

func getListOfWords() []string {
	file, err := os.Open("word-lists/english.txt")

	if err != nil {
		fmt.Println(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words
}
