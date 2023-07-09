package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "localhost:4000"

type errMsg struct{ err error }
type httpResponseMsg string
type statusMessage int

func (e errMsg) Error() string {
	return e.err.Error()
}

func checkserver() tea.Msg {
	c := http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)
	if err != nil {
		return errMsg{err: err}
	}

	return statusMessage(res.StatusCode)

}

type Model struct {
	requestURL string
	response   string
	status     statusMessage
	error      error
}

func initialModel() Model {
	return Model{
		requestURL: "localhost:4000/v1/api",
		response:   "",
	}
}

func (m Model) Init() tea.Cmd {
	return checkserver
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMessage:
		m.status = msg
		return m, tea.Quit

	case errMsg:
		m.error = msg
		return m, tea.Quit

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Model) View() string {
	if m.error != nil {
		return fmt.Sprintf("Could not fetch: %v\n", m.error)
	}

	s := fmt.Sprintf("Checking url... %s", url)

	if m.status > 0 {
		s += fmt.Sprintf("%d %s!", m.status, http.StatusText(int(m.status)))
	}

	return "\n" + s + "\n\n"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("BlazeMorph terminated due to an error %v", err)
		os.Exit(1)
	}
}
