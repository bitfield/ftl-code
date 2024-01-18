package creditcard

import "errors"

type card struct {
	number string
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("number must not be empty")
	}
	return card{number}, nil
}

func (c card) Number() string {
	return c.number
}
