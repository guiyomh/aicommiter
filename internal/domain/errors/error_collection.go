package errors

import "strings"

// ErrorCollection is a collection of errors.
type ErrorCollection struct {
	errors []error
}

// Add adds an error to the collection.
func (e *ErrorCollection) Add(err error) {
	if err != nil {
		e.errors = append(e.errors, err)
	}
}

// Error returns a string representation of the collection.
func (e *ErrorCollection) Error() string {
	if len(e.errors) == 0 {
		return ""
	}

	var sb strings.Builder
	for _, err := range e.errors {
		sb.WriteString(err.Error())
		sb.WriteString("\n")
	}
	return sb.String()
}

// HasErrors vérifie si la collection contient des erreurs.
func (e *ErrorCollection) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *ErrorCollection) Errors() []error {
	return e.errors
}
