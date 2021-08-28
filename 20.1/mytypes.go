package mytypes

// MyInt is a custom version of the `int` type.
type MyInt int

// Twice multiplies its receiver by 2 and returns
// the result.
func (i MyInt) Twice() MyInt {
	return i * 2
}
