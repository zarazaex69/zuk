package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/zarazaex69/zuk/internal/search"
)

type state int

const (
	stateInput state = iota
	stateLoading
	stateResults
)

type Model struct {
	state       state
	query       string
	cursor      int
	results     []search.Result
	selectedIdx int
	err         error
}

func NewModel() Model {
	return Model{
		state: stateInput,
		query: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
