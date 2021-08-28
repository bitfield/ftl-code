package mytypes

// MyString is a custom version of the `string` type.
type MyString string

// Len returns the length of the string.
func (s MyString) Len() int {
	return len(s)
}
