package doctor

import (
	"os/exec"

	"github.com/guiyomh/aicommitter/internal/domain/errors"
)

type DoctorUsecase struct {
}

func New() *DoctorUsecase {
	return &DoctorUsecase{}
}

func (usecase *DoctorUsecase) Execute() error {
	errors := &errors.ErrorCollection{}

	executables := []string{"git", "ollama"}

	for _, execName := range executables {
		err := usecase.checkExec(execName)
		if err != nil {
			errors.Add(err)
		}
	}

	if errors.HasErrors() {
		return errors
	}
	return nil
}

func (usecase *DoctorUsecase) checkExec(execName string) error {
	_, err := exec.LookPath(execName)
	if err != nil {
		return errors.NewExecNotFoundError(execName)
	}
	return nil
}
