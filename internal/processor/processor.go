package processor

import (
	"fmt"
	"github.com/xhos/keisai.md/internal/parser"
	"github.com/xhos/keisai.md/internal/templates"
	"os"
	"path/filepath"
	"strings"
)

// Config holds site generation configuration
type Config struct {
	InputDir     string // directory containing markdown files
	OutputDir    string // directory where HTML files will be written
	DefaultTheme string // default theme to use when none is specified
}

// GenerateSite processes all markdown files in the input directory
func GenerateSite(config Config) error {
	if err := os.MkdirAll(config.OutputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	var markdownFiles []string
	err := filepath.Walk(config.InputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip directories and non-markdown files
		if info.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".md") {
			return nil
		}

		published, err := parser.IsPublished(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Warning: Failed to check if %s is published: %v\n", path, err)
			return nil
		}

		if !published {
			fmt.Printf("Skipping unpublished file: %s\n", path)
			return nil
		}

		markdownFiles = append(markdownFiles, path)
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to traverse input directory: %w", err)
	}

	for _, mdFile := range markdownFiles {
		pageData, err := parser.GetPageData(mdFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error processing %s: %v\n", mdFile, err)
			continue
		}

		theme := pageData.Theme
		if theme == "" {
			theme = config.DefaultTheme
		}

		// apply template
		htmlBuffer, err := templates.WrapInTemplate(pageData, theme)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error applying template to %s: %v\n", mdFile, err)
			continue
		}

		// determine output file path
		relPath, err := filepath.Rel(config.InputDir, mdFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error determining relative path for %s: %v\n", mdFile, err)
			continue
		}
		outputPath := filepath.Join(
			config.OutputDir,
			strings.TrimSuffix(relPath, filepath.Ext(relPath))+".html",
		)

		// create output subdirectories if needed
		outputDir := filepath.Dir(outputPath)
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating output directory for %s: %v\n", outputPath, err)
			continue
		}

		// write the file
		if err := os.WriteFile(outputPath, htmlBuffer.Bytes(), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", outputPath, err)
			continue
		}

		fmt.Printf("Generated %s from %s\n", outputPath, mdFile)
	}

	return nil
}
