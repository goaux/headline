package headline_test

import (
	"fmt"

	"github.com/goaux/headline"
)

func Example() {
	text := "   \n\nFirst line\nSecond line"
	firstLine := headline.Get(text)
	fmt.Println(firstLine)
	// output:
	// First line
}
