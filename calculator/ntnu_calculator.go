package calculator

import (
	"sync"
)

type NTNUCalculator struct {
	Courses map[string]*Course
}

func NewNTNUCalculator(c map[string]*Course) UniversityCalculator {
	return &NTNUCalculator{
		Courses: c,
	}
}

// Rules for NTNU grading:
//
// --- Grade is converted to number value representing the grade (A-E === 5-1).
// --- Grade is multiplied by the points of the course (weighted result).
// --- All weighted grades are then divided by the total amount of points among all courses.
func (n *NTNUCalculator) CalculateGPA() float64 {
	if len(n.Courses) == 0 {
		return 0
	}

	var sum float64
	var totalPoints float64
	var wg sync.WaitGroup

	for _, c := range n.Courses {
		wg.Add(1)
		go func(grade string, points float64) {
			defer wg.Done()
			totalPoints += points
			sum += GradeMap[grade] * points
		}(c.Grade, c.Points)
	}

	wg.Wait()
	return RoundFloat((sum / totalPoints), 1)
}
