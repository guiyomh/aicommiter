package errors

type ExecNotFoundError struct {
	ExecName string
}

func (e ExecNotFoundError) Error() string {
	return "exec not found: " + e.ExecName
}

func NewExecNotFoundError(execName string) ExecNotFoundError {
	return ExecNotFoundError{ExecName: execName}
}
