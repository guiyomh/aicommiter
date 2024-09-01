package promptbuilder

import (
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/rs/zerolog/log"
)

const (
	identity = "You are to act as the author of a commit message in git."
)

type PromptOption func(dp DefaultPromptBuilder) DefaultPromptBuilder

type DefaultPromptBuilder struct {
	spec     CommitSpecification
	useEmoji bool
}

func NewDefaultPromptBuilder(options ...PromptOption) *DefaultPromptBuilder {
	dp := DefaultPromptBuilder{
		spec:     Conventional,
		useEmoji: false,
	}
	for _, option := range options {
		dp = option(dp)
	}
	log.Debug().Str("spec", string(dp.spec)).Bool("useEmoji", dp.useEmoji).Msg("PromptBuilder created")
	return &dp
}

func (b *DefaultPromptBuilder) Build(p domain.Prompt) string {
	prompt := identity + "Your mission is to create clean and comprehensive commit messages as per the "
	switch b.spec {
	case GitMoji:
		prompt += "GitMoji specification"
	case Conventional:
		prompt += "conventional commit convention"
	default:
		prompt += "conventional commit convention"
	}
	prompt += "and explain WHAT were the changes and mainly WHY the changes were done.\n"
	prompt += "I'll send you an output of 'git diff --staged' command, and you are to convert it into a commit message.\n"

	if b.useEmoji {
		prompt += "\nUse GitMoji convention to preface the commit." +
			"Here are some help to choose the right emoji (emoji, description): \n" +
			"🐛, Fix a bug; \n" +
			"✨, Introduce new features; \n" +
			"📝, Add or update documentation; \n" +
			"🚀, Deploy stuff; \n" +
			"✅, Add, update, or pass tests; \n" +
			"♻️, Refactor code; \n" +
			"⬆️, Upgrade dependencies; \n" +
			"🔧, Add or update configuration files; \n" +
			"🌐, Internationalization and localization; \n" +
			"💡, Add or update comments in source code; \n" +
			"🎨, Improve structure / format of the code; \n" +
			"⚡️, Improve performance; \n" +
			"🔥, Remove code or files; \n" +
			"🚑️, Critical hotfix; \n" +
			"💄, Add or update the UI and style files; \n" +
			"🎉, Begin a project; \n" +
			"🔒️, Fix security issues; \n" +
			"🔐, Add or update secrets; \n" +
			"🔖, Release / Version tags; \n" +
			"🚨, Fix compiler / linter warnings; \n" +
			"🚧, Work in progress; \n" +
			"💚, Fix CI Build; \n" +
			"⬇️, Downgrade dependencies; \n" +
			"📌, Pin dependencies to specific versions; \n" +
			"👷, Add or update CI build system; \n" +
			"📈, Add or update analytics or track code; \n" +
			"➕, Add a dependency; \n" +
			"➖, Remove a dependency; \n" +
			"🔨, Add or update development scripts; \n" +
			"✏️, Fix typos; \n" +
			"💩, Write bad code that needs to be improved; \n" +
			"⏪️, Revert changes; \n" +
			"🔀, Merge branches; \n" +
			"📦️, Add or update compiled files or packages; \n" +
			"👽️, Update code due to external API changes; \n" +
			"🚚, Move or rename resources (e.g.: files, paths, routes); \n" +
			"📄, Add or update license; \n" +
			"💥, Introduce breaking changes; \n" +
			"🍱, Add or update assets; \n" +
			"♿️, Improve accessibility; \n" +
			"🍻, Write code drunkenly; \n" +
			"💬, Add or update text and literals; \n" +
			"🗃️, Perform database related changes; \n" +
			"🔊, Add or update logs; \n" +
			"🔇, Remove logs; \n" +
			"👥, Add or update contributor(s); \n" +
			"🚸, Improve user experience / usability; \n" +
			"🏗️, Make architectural changes; \n" +
			"📱, Work on responsive design; \n" +
			"🤡, Mock things; \n" +
			"🥚, Add or update an easter egg; \n" +
			"🙈, Add or update a .gitignore file; \n" +
			"📸, Add or update snapshots; \n" +
			"⚗️, Perform experiments; \n" +
			"🔍️, Improve SEO; \n" +
			"🏷️, Add or update types; \n" +
			"🌱, Add or update seed files; \n" +
			"🚩, Add, update, or remove feature flags; \n" +
			"🥅, Catch errors; \n" +
			"💫, Add or update animations and transitions; \n" +
			"🗑️, Deprecate code that needs to be cleaned up; \n" +
			"🛂, Work on code related to authorization, roles and permissions; \n" +
			"🩹, Simple fix for a non-critical issue; \n" +
			"🧐, Data exploration/inspection; \n" +
			"⚰️, Remove dead code; \n" +
			"🧪, Add a failing test; \n" +
			"👔, Add or update business logic; \n" +
			"🩺, Add or update healthcheck; \n" +
			"🧱, Infrastructure related changes; \n" +
			"🧑‍💻, Improve developer experience; \n" +
			"💸, Add sponsorships or money related infrastructure; \n" +
			"🧵, Add or update code related to multithreading or concurrency; \n" +
			"🦺, Add or update code related to validation; \n"
	} else {
		prompt += "Do not preface the commit with anything. Conventional commit keywords:" +
			"fix, feat, build, chore, ci, docs, style, refactor, perf, test."
	}

	prompt += "\n Add a short description of WHY the changes are done after the commit message." +
		"Don't start it with \"This commit\", just describe the changes." +
		"Use the present tense. Lines must not be longer than 74 characters."
	if p.CommitType != "" {
		prompt += "\nUse the type: " + p.CommitType
	}
	if p.Scope != "" {
		prompt += "\nUse the scope: " + p.Scope
	}
	if p.IssueNumber != "" {
		prompt += "\nSpecify the issue: #" + p.IssueNumber
	}
	if p.Language != "" {
		prompt += "\nUse  " + p.Language + " language for the commit message"
	}

	// prompt += "\ngit diff --staged output:\n```\n"
	// prompt += p.Diff.Content
	// prompt += "\n```"

	log.Trace().Str("prompt", prompt).Msg("Prompt created")

	return prompt
}
