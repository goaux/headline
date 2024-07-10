// Package headline provides functionality to extract the first non-empty line from a string.
//
// This package is useful for processing text data where the first meaningful line
// needs to be extracted, such as in configuration files, headers, or any text
// where leading empty lines or whitespace should be ignored.
//
// The package contains a single function, Get, which efficiently processes
// the input string and returns the first line containing non-whitespace characters.
// The returned string does not include any newline characters.
//
// Usage:
//
//	import "github.com/goaux/headline"
//
//	text := "   \n\nFirst line\nSecond line"
//	firstLine := headline.Get(text)
//	// firstLine will be "First line"
package headline

import "strings"

// Get retrieves the first non-empty line from a string.
// It returns a string without any newline characters.
func Get(s string) string {
	for {
		// If the string is empty, return it as is
		if s == "" {
			return s
		}

		// Find the index of the first newline character
		i := strings.IndexByte(s, '\n')

		// If no newline is found, return the entire remaining string
		// Note: This case also ensures no newline is included
		if i == -1 {
			return s
		}

		// Extract the first line (excluding the newline) and trim whitespace
		v := s[:i]
		if strings.TrimSpace(v) != "" {
			// If a non-empty line is found, return it
			// Note: The newline is not included in the returned string
			return v
		}

		// If the line is empty, move to the next line
		s = s[i+1:]
	}
}
