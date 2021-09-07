// Package calculator provides a library for
// simple calculations in Go.
package calculator

import (
	"errors"
)

func AddMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, n := range inputs[1:] {
		result += n
	}
	return result
}

func MultiplyMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, n := range inputs[1:] {
		result *= n
	}
	return result
}

func SubtractMany(inputs ...float64) float64 {
	if len(inputs) == 0 {
		return 0
	}
	result := inputs[0]
	for _, n := range inputs[1:] {
		result -= n
	}
	return result
}

func DivideMany(inputs ...float64) (float64, error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	result := inputs[0]
	for _, n := range inputs[1:] {
		if n == 0 {
			return 0, errors.New("division by zero not allowed")
		}
		result /= n
	}
	return result, nil
}
