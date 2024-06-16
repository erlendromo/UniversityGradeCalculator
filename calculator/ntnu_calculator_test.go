package calculator

import "testing"

type NTNUCalculatorTestCase struct {
	Courses  map[string]*Course
	Expected float64
}

var testCases = []NTNUCalculatorTestCase{
	{
		Courses: map[string]*Course{
			"1": {Grade: "A", Points: 10},
			"2": {Grade: "B", Points: 5},
		},
		Expected: 4.7,
	},
	{
		Courses: map[string]*Course{
			"1": {Grade: "A", Points: 10},
			"2": {Grade: "A", Points: 10},
		},
		Expected: 5,
	},
	{
		Courses: map[string]*Course{
			"1": {Grade: "A", Points: 10},
			"2": {Grade: "D", Points: 10},
		},
		Expected: 3.5,
	},
}

func TestNTNUCaclulateGPA(t *testing.T) {
	// Create a new NTNU calculator with a map of courses.
	for _, tc := range testCases {
		calc := NewNTNUCalculator(tc.Courses)

		// Calculate the GPA.
		got := calc.CalculateGPA()

		// Check if the result is as expected.
		if got != tc.Expected {
			t.Errorf("Expected %v, got %v", tc.Expected, got)
		}
	}
}
