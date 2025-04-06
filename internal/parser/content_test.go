package parser

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

func TestGetFileName(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
	}{
		{
			name:     "simple filename",
			filePath: "/path/to/myFile.txt",
			want:     "myfile",
		},
		{
			name:     "filename with spaces",
			filePath: "/path/to/my File with spaces.txt",
			want:     "my-file-with-spaces",
		},
		{
			name:     "filename with linux special characters",
			filePath: "/path/to/my !'\"\\|><file.txt",
			want:     "my-file",
		},
		{
			name:     "filename with multiple dots",
			filePath: "/path/to/my.file.name.with.dots.txt",
			want:     "my-file-name-with-dots",
		},
		{
			name:     "filename with leading/trailing spaces",
			filePath: "/path/to/  my file  .txt",
			want:     "my-file",
		},
		{
			name:     "filename with leading/trailing special characters",
			filePath: "/path/to/!@#my$file%.txt",
			want:     "my-file",
		},
		{
			name:     "filename with consecutive special characters",
			filePath: "/path/to/my---file.txt",
			want:     "my-file",
		},
		{
			name:     "filename with no extension",
			filePath: "/path/to/myFile",
			want:     "myfile",
		},
		{
			name:     "filename with only extension",
			filePath: "/path/to/.txt",
			want:     "",
		},
		{
			name:     "empty file path",
			filePath: "",
			want:     "",
		},
		{
			name:     "unicode test",
			filePath: "/path/to/你好世界.txt",
			want:     "你好世界",
		},
		{
			name:     "unicode with symbols",
			filePath: "/path/to/你好@#$世界%.txt",
			want:     "你好-世界",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFileName(tt.filePath)
			if got != tt.want {
				t.Errorf("getSafeFileName(%q) = %q, want %q", tt.filePath, got, tt.want)
			}
		})
	}
}
