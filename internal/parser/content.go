package parser

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"html/template"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type PageData struct {
	Title   string
	Path    string
	Theme   string
	Content template.HTML
}

// path -> PageData
func GetPageData(filePath string) (PageData, error) {
	// Convert markdown to HTML
	htmlBuf, err := convertToHTML(filePath)
	if err != nil {
		return PageData{}, err
	}

	// Get theme, handling potential error
	theme, err := ParseTheme(filePath)
	if err != nil {
		return PageData{}, err
	}

	// Create and return PageData
	return PageData{
		Title:   getFileName(filePath),
		Path:    filePath,
		Theme:   theme,
		Content: template.HTML(htmlBuf.String()),
	}, nil
}

// path -> buffer of html
func convertToHTML(filePath string) (*bytes.Buffer, error) {
	// read file
	source, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	// parse markdown
	var buf bytes.Buffer
	if err = goldmark.Convert(source, &buf); err != nil {
		return nil, fmt.Errorf("failed to convert markdown: %w", err)
	}

	return &buf, nil
}

// path -> name string
func getFileName(filePath string) string {
	baseName := filepath.Base(filePath)
	fileName := strings.TrimSuffix(baseName, filepath.Ext(baseName))

	var safeFileName strings.Builder
	lastCharWasHyphen := false

	for _, r := range fileName {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			safeFileName.WriteRune(unicode.ToLower(r))
			lastCharWasHyphen = false
		} else {
			if !lastCharWasHyphen && safeFileName.Len() > 0 {
				safeFileName.WriteRune('-')
				lastCharWasHyphen = true
			}
		}
	}

	result := strings.Trim(safeFileName.String(), "-")
	return result
}
