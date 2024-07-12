# headline
Package headline provides functionality to extract the first non-empty line from a string.

[![Go Reference](https://pkg.go.dev/badge/github.com/goaux/headline.svg)](https://pkg.go.dev/github.com/goaux/headline)
[![Go Report Card](https://goreportcard.com/badge/github.com/goaux/headline)](https://goreportcard.com/report/github.com/goaux/headline)

This package is useful for processing text data where the first meaningful line
needs to be extracted, such as in configuration files, headers, or any text
where leading empty lines or whitespace should be ignored.

The package contains a single function, Get, which efficiently processes
the input string and returns the first line containing non-whitespace characters.
The returned string does not include any newline characters.

Usage:

    import "github.com/goaux/headline"

    text := "   \n\nFirst line\nSecond line"
    firstLine := headline.Get(text)
    // firstLine will be "First line"
