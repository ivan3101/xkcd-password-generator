package main

import (
	"fmt"
	"io/fs"
	"os"
)

func main() {
	filesystem := os.DirFS("word-lists")
	fileContent, err := fs.ReadFile(filesystem, "english.txt")

	if err != nil {
		fmt.Println(err)
	}

	words := string(fileContent)

	fmt.Println(words)
}
