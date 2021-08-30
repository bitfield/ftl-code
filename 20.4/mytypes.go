package mytypes

import "strings"

// MyBuilder embeds strings.Builder.
type MyBuilder struct {
	strings.Builder
}
