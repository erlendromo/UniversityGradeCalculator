package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"github.com/erlendromo/UniversityGradeCalculator/calculator"
)

func guard(err error) {
	if err != nil {
		panic(err)
	}
}

func StartServer() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /calculator/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./web/index.html"))
		guard(tmpl.Execute(w, nil))
	})

	mux.HandleFunc("POST /add-course/", func(w http.ResponseWriter, r *http.Request) {
		code := r.PostFormValue("code")
		name := r.PostFormValue("name")
		grade := r.PostFormValue("grade")
		pointsStr := r.PostFormValue("points")

		validStrings := []calculator.ValidString{
			{Key: "code", Value: code, Length: 20},
			{Key: "name", Value: name, Length: 50},
			{Key: "grade", Value: grade, Length: 1},
			{Key: "points", Value: pointsStr, Length: 4},
		}
		errs := calculator.IsValidStrings(validStrings)

		grade = strings.ToUpper(grade)
		if !calculator.IsValidGrade(grade) {
			errs["grade"] = fmt.Sprintf("Invalid input: %s", grade)
		}

		points, err := strconv.ParseFloat(r.PostFormValue("points"), 64)
		if err != nil {
			errs["points"] = fmt.Sprintf("Invalid input: %s", pointsStr)
		} else if points == 0 {
			errs["points"] = fmt.Sprintf("Invalid input: %.1f", points)
		}

		if len(errs) > 0 {
			resp := calculator.NewErrorResponse(errs)
			log.Println(resp)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if calculator.CourseExists(code) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		calculator.AppendCourse(&calculator.Course{
			Code:   code,
			Name:   name,
			Grade:  grade,
			Points: points,
		})

		htmlStr := fmt.Sprintf("<div class='added'><p>Code: %s</p><p>Name: %s</p><p>Grade: %v</p><p>Points: %v</p><div>", code, name, grade, points)
		tmpl := template.Must(template.New("t").Parse(htmlStr))
		guard(tmpl.Execute(w, nil))
	})

	mux.HandleFunc("POST /calc-grades/", func(w http.ResponseWriter, r *http.Request) {
		university := r.PostFormValue("university")

		c, ok := calculator.GetCalculator(university)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		grade := c.CalculateGPA()
		if grade == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		letter := calculator.LetterEquivalent(grade)

		htmlStr := fmt.Sprintf("<p>Grade: %.1f ----- Letter: %s</p>", grade, letter)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	})

	log.Println("Starting server on port 8080...")
	guard(http.ListenAndServe(":8080", mux))
}
