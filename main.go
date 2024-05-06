package main

import (
	"strings"
	"terminal-gpt/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

const TRUE bool = true
const FALSE bool = false

var Program *tea.Program

func main() {
	w, _, _ := term.GetSize(0)
	inw := 60
	if w-4 < inw {
		inw = w - 4
	}
	placeholder := "Describe the command you want to run"
	if len(placeholder) < inw {
		dif := inw - len(placeholder) + 1
		placeholder = placeholder + strings.Repeat(" ", dif)
	}
	var focus bool = TRUE
	in := components.GenerateTextInput(&components.TextboxOptions{
		Placeholder: &placeholder,
		Focus:       &focus,
	})

	model := Model{
		Loading:        false,
		Input:          in,
		Actions:        []string{"Copy to clipboard", "Edit Command", "Ask something else", "Quit"},
		spinner:        components.GenerateSpinner(nil),
		selectedAction: 0,
		bannerText:     "",
		bannerStyle:    lipgloss.NewStyle(),
		OllamaModel:    Gemma7b,
	}
	Program = tea.NewProgram(model, tea.WithAltScreen())
	_, e := Program.Run()
	if e != nil {
		panic(e)
	}
}

func s() lipgloss.Style {
	return lipgloss.NewStyle()
}
