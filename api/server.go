package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		tmp := template.Must(template.ParseFiles("./web/index.html"))
		guard(tmp.Execute(w, nil))
	})

	mux.HandleFunc("POST /add-course/", func(w http.ResponseWriter, r *http.Request) {
		code := r.PostFormValue("code")
		name := r.PostFormValue("name")
		grade := r.PostFormValue("grade")
		points, err := strconv.ParseFloat(r.PostFormValue("points"), 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		calculator.AppendCourse(&calculator.Course{
			Code:   code,
			Name:   name,
			Grade:  grade,
			Points: points,
		})

		htmlStr := fmt.Sprintf("<div class='added'><p>Code: %s</p><p>Name: %s</p><p>Grade: %v</p><p>Points: %v</p><div>", code, name, grade, points)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	})

	mux.HandleFunc("POST /calc-grades/", func(w http.ResponseWriter, r *http.Request) {
		university := r.PostFormValue("university")

		c, ok := calculator.GetCalculator(university)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		grade := c.CalculateGPA()
		letter := calculator.LetterEquivalent(grade)

		htmlStr := fmt.Sprintf("<p>Grade: %.1f ----- Letter: %s</p>", grade, letter)
		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	})

	log.Println("Starting server on port 8080...")
	guard(http.ListenAndServe(":8080", mux))
}
