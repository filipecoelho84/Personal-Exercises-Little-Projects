package calculator_test

import (
	"calculator"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCaseSub struct {
		a, b float64
		want float64
	}

	testCaseSubt := []testCaseSub{
		{a: 0, b: 5, want: -5},
		{a: 5, b: 2.5, want: 2.5},
		{a: 2.5, b: 0, want: 2.5},
	}

	for _, tc1 := range testCaseSubt {
		got := calculator.Subtract(tc1.a, tc1.b)
		if tc1.want != got {
			t.Errorf("Sub(%f, %f): want %f, got %f", tc1.a, tc1.b, tc1.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCaseMul struct {
		a, b float64
		want float64
	}

	testCaseMult := []testCaseMul{
		{a: 1, b: 3, want: 3},
		{a: 0, b: 99, want: 0},
		{a: 10, b: 20, want: 200},
	}

	for _, tcm := range testCaseMult {
		got := calculator.Multiply(tcm.a, tcm.b)
		if tcm.want != got {
			t.Errorf("Mult(%f, %f): want %f, got %f", tcm.a, tcm.b, tcm.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func closeEnough(a, b, tolelance float64) bool {
	return math.Abs(a-b) <= tolelance
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Divide(1, 0): want error for invalid input, got nil")
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a    float64
		want float64
	}

	testCases := []testCase{
		{a: 2, want: 1.414213},
		{a: 3, want: 1.732050},
		{a: 9, want: 3},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Errorf("Sqrt(%f): is an invalid input!", tc.a)
		}
		if !closeEnoughSqrt(tc.want, got, 0.001) {
			t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		}
	}
}

func closeEnoughSqrt(a, b, tolelance float64) bool {
	return math.Abs(a-b) <= tolelance
}
