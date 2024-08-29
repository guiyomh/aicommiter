package promptbuilder

// WithEmoji is an option to use emoji in the prompt.
func WithEmoji() PromptOption {
	return func(dp DefaultPromptBuilder) DefaultPromptBuilder {
		dp.useEmoji = true
		return dp
	}
}

// WithSpecification is an option to use a specific commit specification in the prompt.
func WithSpecification(spec CommitSpecification) PromptOption {
	return func(dp DefaultPromptBuilder) DefaultPromptBuilder {
		dp.spec = spec
		return dp
	}
}
