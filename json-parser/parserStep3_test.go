// +build step3

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
		// Define test cases for step3
		{"step3", "valid.json", true},
		{"step3", "invalid.json", false},
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
				t.Errorf("Got %q, which is %t.", msg, tc.validity)
			}
		})
	}
}
