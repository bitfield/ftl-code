package mytypes

type MyInt int

func (input *MyInt) Double() {
	*input *= 2
}
