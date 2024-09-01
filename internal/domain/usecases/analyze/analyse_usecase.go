package analyze

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain/errors"
	"github.com/guiyomh/aicommitter/internal/domain/services/commitmessage"
	"github.com/guiyomh/aicommitter/internal/domain/services/gitdiff"
	"github.com/guiyomh/aicommitter/internal/domain/services/promptbuilder"
)

type AnalyzeUsecase struct {
}

func New() *AnalyzeUsecase {
	return &AnalyzeUsecase{}
}

func (a AnalyzeUsecase) Analyze(ctx context.Context, input Input) error {

	pb := promptbuilder.NewDefaultPromptBuilder(
		promptbuilder.WithSpecification(input.CommitSpecification),
	)

	adapter, err := commitmessage.NewAdapter(ctx, input.AdapterType, pb, input.ApiKey)
	if err != nil {
		return errors.NewAnalyzeError("Error while creating adapter: " + err.Error())
	}

	service := commitmessage.New(adapter)

	diffGenerator := gitdiff.New()
	diff, err := diffGenerator.GenerateDiff()
	if err != nil {
		return errors.NewAnalyzeError("Error while generating diff: " + err.Error())
	}

	commitMessage, err := service.CreateCommitMessage(
		ctx,
		diff,
		input.CommitOptions...,
	)
	if err != nil {
		return errors.NewAnalyzeError("Error while generating commit message: " + err.Error())
	}
	fmt.Printf("Generated Commit Message:\n\n%s\n%s\n%s\n", commitMessage.Header, commitMessage.Body, commitMessage.Footer)
	return nil
}
