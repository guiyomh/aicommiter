package promptbuilder

type Option func(dp DefaultPromptBuilder) DefaultPromptBuilder

// WithEmoji is an option to use emoji in the prompt.
func WithEmoji() Option {
	return func(dp DefaultPromptBuilder) DefaultPromptBuilder {
		dp.useEmoji = true
		return dp
	}
}

// WithSpecification is an option to use a specific commit specification in the prompt.
func WithSpecification(spec CommitSpecification) Option {
	return func(dp DefaultPromptBuilder) DefaultPromptBuilder {
		dp.spec = spec
		return dp
	}
}
