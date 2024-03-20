//go:build step2
// +build step2

package jsonparser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseJSONFiles(t *testing.T) {
	testCases := []struct {
		step     string
		fileName string
		validity bool
	}{
		// Define test cases for step2
		{"step2", "valid.json", true},
		{"step2", "valid1.json", true},
		{"step2", "valid2.json", true},
		{"step2", "valid3.json", true},
		{"step2", "invalid.json", false},
		{"step2", "invalid1.json", false},
		{"step2", "invalid2.json", false},
		{"step2", "invalid3.json", false},
	}

	for _, tc := range testCases {
		t.Run(tc.step+"/"+tc.fileName, func(t *testing.T) {
			filePath := filepath.Join("json-files", tc.step, tc.fileName)
			file, err := os.Open(filePath)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
			}
			defer file.Close()

			json_data, err := Lexer(file)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			isValid, msg := Parse(json_data)

			if !tc.validity && isValid {
				t.Errorf(msg)
			}
		})
	}
}
