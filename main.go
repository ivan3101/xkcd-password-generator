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
	var indexesUsed = make(map[int64]string)
	var password string

	for i := 1; i <= totalOfWordsForPassword; i++ {
		wasIndexAlreadyUsed := true
		var randomIndex int64

		for wasIndexAlreadyUsed {
			randomIndex = getRandomIndex(totalWords)

			wasIndexAlreadyUsed = checkIfIndexWasAlreadyUsed(randomIndex, indexesUsed)
		}

		indexesUsed[randomIndex] = words[randomIndex]

		if len(password) == 0 {
			password = words[randomIndex]
			continue
		}

		password += " " + words[randomIndex]
	}

	return password
}

func getRandomIndex(totalWords *big.Int) int64 {
	randomIndex, err := rand.Int(rand.Reader, totalWords)

	if err != nil {
		fmt.Println(err)
	}

	return randomIndex.Int64()
}

func checkIfIndexWasAlreadyUsed(index int64, indexesUsed map[int64]string) bool {
	_, exists := indexesUsed[index]

	return exists
}
