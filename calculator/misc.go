package calculator

import (
	"fmt"
	"strings"
)

type ErrorResponse struct {
	Errors map[string]string `form:"errors"`
}

func NewErrorResponse(errs map[string]string) *ErrorResponse {
	return &ErrorResponse{Errors: errs}
}

type ValidString struct {
	Key    string
	Value  string
	Length int
}

func IsEmptyString(s string) bool {
	return strings.TrimSpace(s) == ""
}

func IsLongerThan(s string, n int) bool {
	return len(s) > n
}

func IsValidString(s string, n int) bool {
	return !IsEmptyString(s) && !IsLongerThan(s, n)
}

func IsValidGrade(grade string) bool {
	return strings.Contains("ABCDE", strings.ToUpper(grade))
}

func IsValidStrings(validStrings []ValidString) map[string]string {
	errs := make(map[string]string)
	for _, vs := range validStrings {
		if !IsValidString(vs.Value, vs.Length) {
			errs[vs.Key] = fmt.Sprintf("Invalid input: %s", vs.Value)
		}
	}
	return errs
}
