package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
)

const totalOfWordsForPassword = 4
const wordsListsDefaultPath = "word-lists"
const filesExtension = "txt"

var languagesAvailable = map[int64]string{
	1: "czech",
	2: "english",
	3: "french",
	4: "italian",
	5: "japanese",
	6: "korean",
	7: "portuguese",
	8: "spanish",
}

func main() {
	lang := selectLanguage()
	wordList := getWordsFile(lang)
	words := getListOfWords(wordList)
	password := generatePassword(words)

	fmt.Println(password)
}

func selectLanguage() string {
	fmt.Println("Languages available")
	fmt.Println("1) Czech (cs)")
	fmt.Println("2) English (en)")
	fmt.Println("3) French (fr)")
	fmt.Println("4) Italian (it)")
	fmt.Println("5) Japanese (jp)")
	fmt.Println("6) Korean (ko)")
	fmt.Println("7) Portuguese (pt)")
	fmt.Println("8) Spanish (es)")

	scanner := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the number of the language that you want to use to generate the password:")

		input, _, err := scanner.ReadLine()

		if err != nil {
			fmt.Println("There was an error. Try again")
			continue
		}

		inputAsNumber, err := strconv.ParseInt(string(input), 10, 64)

		if err != nil {
			fmt.Println("You must enter a number. Try again")
			continue
		}

		lang, isAValidOption := languagesAvailable[inputAsNumber]

		if !isAValidOption {
			fmt.Println("You must enter a valid language")
			continue
		}

		return lang
	}
}

func getWordsFile(lang string) *os.File {
	file, err := os.OpenFile(wordsListsDefaultPath+"/"+lang+"."+filesExtension, os.O_RDONLY, io.SeekStart)

	if err != nil {
		fmt.Println("Error opening the file")
		os.Exit(2)
	}

	return file
}

func getListOfWords(wordList *os.File) []string {
	defer func(file *os.File) {
		_ = file.Close()
	}(wordList)

	scanner := bufio.NewScanner(wordList)

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
