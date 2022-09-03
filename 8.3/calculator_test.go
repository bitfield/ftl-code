package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{2, 2}, want: 4},
		{inputs: []float64{1, 2, 3}, want: 6},
	}
	for _, tc := range testCases {
		got := calculator.AddMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("AddMany(%v): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestSubtractMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{2, 2}, want: 0},
		{inputs: []float64{1, 2, 3}, want: -4},
	}
	for _, tc := range testCases {
		got := calculator.SubtractMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("SubtractMany(%v): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestMultiplyMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{2, 2}, want: 4},
		{inputs: []float64{1, 2, 3, 4}, want: 24},
	}
	for _, tc := range testCases {
		got := calculator.MultiplyMany(tc.inputs...)
		if tc.want != got {
			t.Errorf("MultiplyMany(%v): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestDivideMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want   float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{2, 2}, want: 1},
		{inputs: []float64{-1, -1}, want: 1},
		{inputs: []float64{100, 5, 2}, want: 10},
	}
	for _, tc := range testCases {
		got, err := calculator.DivideMany(tc.inputs...)
		if err != nil {
			t.Fatalf("DivideMany(%v): want no error for valid input, got %v", tc.inputs, err)
		}
		if tc.want != got {
			t.Errorf("DivideMany(%v): want %f, got %f",
				tc.inputs, tc.want, got)
		}
	}
}

func TestDivideManyInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.DivideMany(10, 5, 0, 7)
	if err == nil {
		t.Error("Divide(10, 5, 0, 7): want error for zero divisor, got nil")
	}
}
