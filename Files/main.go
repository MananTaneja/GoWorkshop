package main

import (
	"fmt"

	"manantaneja.com/go/files/fileutils"
)

func main() {
	content, err := fileutils.ReadTextFile("sample.txt")

	if err != nil {
		panic("Unable to read file contents")
	}

	fmt.Printf("content: %s\n", content)

	writeContent := fmt.Sprintf("Original: %v\nUpdated: %v %v", content, content, content)

	fileutils.WriteTextFile("output.txt", writeContent)

}
