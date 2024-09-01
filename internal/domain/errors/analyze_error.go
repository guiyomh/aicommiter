package errors

type AnalyzeError struct {
	Message string
}

func (e AnalyzeError) Error() string {
	return e.Message
}

func NewAnalyzeError(message string) AnalyzeError {
	return AnalyzeError{Message: message}
}
