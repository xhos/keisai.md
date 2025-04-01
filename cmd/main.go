package main

import (
	"fmt"
	"github.com/xhos/keisai.md/internal/markdown"
	"os"
)

func main() {
	filePath := ".github/readme.md"

	result, err := markdown.ConvertToHTML(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result.String())
}
