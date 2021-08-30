package mytypes

import "strings"

// StringUppercaser wraps strings.Builder.
type StringUppercaser struct {
	Contents strings.Builder
}

func (su StringUppercaser) ToUpper() string {
	return strings.ToUpper(su.Contents.String())
}
