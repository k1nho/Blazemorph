package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "localhost:4000"

type Model struct {
	requestURL string
	response   interface{}
}

func initialModel() Model {
	return Model{
		requestURL: "localhost:4000/v1/api",
		response:   "hello from endpoint",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Model) View() string {
	s := "BlazeMorph"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("BlazeMorph terminated due to an error %v", err)
		os.Exit(1)
	}
}
