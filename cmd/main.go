package main

import (
	"flag"
	"fmt"
	"github.com/xhos/keisai.md/internal/processor"
	"os"
	"path/filepath"
)

func main() {
	inputDir := flag.String("input", "", "Directory containing markdown files")
	outputDir := flag.String("output", "public", "Directory where HTML files will be written")
	theme := flag.String("theme", "default", "Default theme to use")

	flag.Parse()

	if *inputDir == "" {
		fmt.Fprintln(os.Stderr, "Error: Input directory must be specified")
		fmt.Fprintln(os.Stderr, "Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	absInputDir, err := filepath.Abs(*inputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving input path: %v\n", err)
		os.Exit(1)
	}

	absOutputDir, err := filepath.Abs(*outputDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving output path: %v\n", err)
		os.Exit(1)
	}

	config := processor.Config{
		InputDir:     absInputDir,
		OutputDir:    absOutputDir,
		DefaultTheme: *theme,
	}

	fmt.Printf("Processing markdown files from: %s\n", config.InputDir)
	fmt.Printf("Writing HTML output to: %s\n", config.OutputDir)
	fmt.Printf("Using default theme: %s\n", config.DefaultTheme)

	if err := processor.GenerateSite(config); err != nil {
		fmt.Fprintf(os.Stderr, "Error generating site: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Site generated successfully!")
}
