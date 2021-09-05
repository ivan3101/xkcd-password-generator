package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

const totalOfWordsForPassword = 4

func main() {
	words := getListOfWords()
	password := generatePassword(words)

	fmt.Println(password)
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

func generatePassword(words []string) string {
	totalWords := big.NewInt(int64(len(words)))
	var password string

	for i := 1; i <= totalOfWordsForPassword; i++ {
		randomIndex, err := rand.Int(rand.Reader, totalWords)

		if err != nil {
			fmt.Println(err)
		}

		if len(password) == 0 {
			password = words[randomIndex.Int64()]
			continue
		}

		password += " " + words[randomIndex.Int64()]
	}

	return password
}
