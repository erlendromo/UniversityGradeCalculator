package calculator

import "testing"

type CalculatorTestCase struct {
	University string
	Expected   bool
}

var testCases2 = []CalculatorTestCase{
	{
		University: "NTNU",
		Expected:   true,
	},
	{
		University: "UiO",
		Expected:   false,
	},
}

func TestGetCalculator(t *testing.T) {
	for _, tc := range testCases2 {
		calc, ok := GetCalculator(tc.University)

		if ok != tc.Expected {
			t.Errorf("Expected %v, got %v", tc.Expected, ok)
		} else if ok && calc == nil {
			t.Errorf("Expected calculator to be non-nil")
		}
	}
}

type LetterEquivalentTestCase struct {
	Grade    float64
	Expected string
}

var testCases3 = []LetterEquivalentTestCase{
	{
		Grade:    5,
		Expected: "A",
	},
	{
		Grade:    4.7,
		Expected: "A",
	},
	{
		Grade:    4.3,
		Expected: "B",
	},
	{
		Grade:    3.5,
		Expected: "B",
	},
	{
		Grade:    2.5,
		Expected: "C",
	},
	{
		Grade:    1.1,
		Expected: "E",
	},
}

func TestLetterEquivalent(t *testing.T) {
	for _, tc := range testCases3 {
		got := LetterEquivalent(tc.Grade)

		if got != tc.Expected {
			t.Errorf("Expected %s, got %s", tc.Expected, got)
		}
	}
}
