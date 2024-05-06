package components

import (
	"terminal-gpt/colors"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type TextboxOptions struct {
	Placeholder *string
	Value       *string
	Focus       *bool
}

func GenerateTextInput(options *TextboxOptions) textinput.Model {
	in := textinput.New()
	inw := 60
	in.Width = inw
	placeholder := options.Placeholder
	if placeholder != nil {
		in.Placeholder = *placeholder
	}

	value := options.Value
	if value != nil {
		in.SetValue(*value)
	}

	if options.Focus != nil && *options.Focus {
		in.Focus()
	}

	in.Prompt = ""
	in.PromptStyle = lipgloss.NewStyle().Foreground(colors.PrimaryContent)
	in.TextStyle = lipgloss.NewStyle().Foreground(colors.Primary).Background(colors.PrimaryContent)
	in.PlaceholderStyle = in.TextStyle
	in.Cursor.Style = lipgloss.NewStyle().Foreground(colors.Secondary).Background(colors.SecondaryContent)
	return in
}
