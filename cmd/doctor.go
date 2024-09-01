package cmd

import (
	"errors"
	"fmt"

	myerrors "github.com/guiyomh/aicommitter/internal/domain/errors"
	"github.com/guiyomh/aicommitter/internal/domain/usecases/doctor"
	"github.com/spf13/cobra"
)

// doctorCmd represents the doctor command
func NewDoctorCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "doctor",
		Short: "Check if the environment is properly set up",
		Run: func(_ *cobra.Command, _ []string) {

			var errorCollection *myerrors.ErrorCollection
			usecase := doctor.New()
			err := usecase.Execute()
			if errors.As(err, &errorCollection) {
				fmt.Println("❌ Doctor found some issues:")
				for _, e := range errorCollection.Errors() {
					fmt.Println(e.Error())
				}
				return
			}
			fmt.Println("✅ Doctor found no issues")
		},
	}
}
