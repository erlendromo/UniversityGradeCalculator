package calculator

import (
	"math"
)

type UniversityCalculator interface {
	CalculateGPA() float64
}

type Course struct {
	Code   string  `form:"code"`
	Name   string  `form:"name"`
	Grade  string  `form:"grade"`
	Points float64 `form:"points"`
}

var (
	courses  map[string]*Course
	GradeMap = map[string]float64{
		"A": 5,
		"B": 4,
		"C": 3,
		"D": 2,
		"E": 1,
	}
)

func init() {
	courses = make(map[string]*Course, 0)
}

func AppendCourse(c *Course) {
	courses[c.Code] = c
}

func GetCalculator(u string) (calc UniversityCalculator, ok bool) {
	ok = true

	switch u {
	case "NTNU":
		calc = NewNTNUCalculator(courses)
	default:
		ok = false
	}

	return calc, ok
}

func LetterEquivalent(grade float64) string {
	for l, g := range GradeMap {
		if g == math.Ceil(grade) {
			return l
		}
	}

	return ""
}
