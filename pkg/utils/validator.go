package utils

import (
	"errors"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/RuhullahReza/Employee-App/app/domain"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	ErrEmptyName    = errors.New("empty name field")
	ErrInvalidEmail = errors.New("invalid email format")
	ErrInvalidName  = errors.New("invalid name format")
)

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func isWhitespace(str string) bool {
	for _, char := range str {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}

func isAlphaAndSpace(str string) bool {
	pattern := "^[a-zA-Z ]+$"

	regex := regexp.MustCompile(pattern)
	return regex.MatchString(str)
}

func sanitizeName(name string) string {
	nameList := strings.Split(name, " ")
	var trimedNameList []string
	for _, s := range nameList {
		if isWhitespace(s) {
			continue
		}

		cap := cases.Title(language.Und, cases.NoLower).String(s)
		trimedNameList = append(trimedNameList, cap)
	}

	return strings.Join(trimedNameList, " ")
}

func ValidateAndSanitizeRequest(req *domain.EmployeeRequest) error {
	firstName := strings.TrimSpace(req.FirstName)
	lastName := strings.TrimSpace(req.LastName)

	if len(firstName) == 0 || len(lastName) == 0 {
		return ErrEmptyName
	}

	if !isAlphaAndSpace(firstName) || !isAlphaAndSpace(lastName) {
		return ErrInvalidName
	}

	if !isValidEmail(req.Email) {
		return ErrInvalidEmail
	}

	req.FirstName = sanitizeName(firstName)
	req.LastName = sanitizeName(lastName)

	return nil
}

func ParseDateString(dateString string) (time.Time, error) {
	layout := "2006-01-02"

	parsedTime, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}

	return parsedTime, nil
}
