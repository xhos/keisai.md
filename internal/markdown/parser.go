package markdown

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"os"
)

func ConvertToHTML(filePath string) (*bytes.Buffer, error) {
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
