package mytypes

import "strings"

// MyBuilder is a custom version of the `strings.Builder` type.
type MyBuilder strings.Builder

// Hello returns the string "Hello, Gophers!"
func (mb MyBuilder) Hello() string {
	return "Hello, Gophers!"
}
