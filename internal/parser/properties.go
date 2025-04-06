package parser

import (
	"fmt"
	"os"
	"strings"
)

// extracts the value of the specified property from a markdown file's frontmatter
func parseProperty(filePath, propertyName string) (string, error) {
	// only process markdown files
	if !strings.HasSuffix(strings.ToLower(filePath), ".md") {
		return "", nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	// split into lines
	fileContent := string(content)
	lines := strings.Split(fileContent, "\n")

	// find frontmatter opening
	if len(lines) < 2 || lines[0] != "---" {
		return "", nil
	}

	// find the frontmatter closing
	var endOfFrontmatter int
	for i := 1; i < len(lines); i++ {
		if lines[i] == "---" {
			endOfFrontmatter = i
			break
		}
	}

	if endOfFrontmatter == 0 {
		return "", nil
	}

	// extract frontmatter lines
	frontmatter := lines[1:endOfFrontmatter]

	// check each line for the property at the root level
	for _, line := range frontmatter {
		// skip if line has leading whitespace (indicating it's nested)
		if len(line) > 0 && (line[0] == ' ' || line[0] == '\t') {
			continue
		}

		// get the key and value
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if key == propertyName {
				return value, nil
			}
		}
	}

	return "", nil
}

// checks if the markdown file is marked as published in its frontmatter
func IsPublished(filePath string) (bool, error) {
	value, err := parseProperty(filePath, "published")
	if err != nil {
		return false, err
	}
	return value == "true", nil
}

// ParseTheme extracts the theme property value from the markdown frontmatter
// If the theme property is empty, it returns "default".
func ParseTheme(filePath string) (string, error) {
	value, err := parseProperty(filePath, "theme")
	if err != nil {
		return "", err
	}
	if value == "" {
		return "default", nil
	}
	return value, nil
}
