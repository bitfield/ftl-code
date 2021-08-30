package mytypes

import "strings"

// MyBuilder is a custom version of the `strings.Builder` type.
type MyBuilder struct {
	Contents strings.Builder
}
