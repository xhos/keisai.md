package markdown

import (
	"testing"
)

// TODO: Write a test for ConvertToHTML
// func TestConvertToHTML(t *testing.T) {}

func TestConvertToHTML_FileNotFound(t *testing.T) {
	_, err := ConvertToHTML("randomfile.md")
	if err == nil {
		t.Error("Expected an error when file doesn't exist, but got nil")
	}
}
