package ui

import (
	"fmt"

	"github.com/stikypiston/fastcards/internal"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	due      []internal.DueRef
	index    int
	showBack bool
	done     bool
}

func NewStudyModelFromRefs(refs []internal.DueRef) *Model {
	return &Model{
		due: refs,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.done {
		return m, tea.Quit
	}

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			m.showBack = true

		case "y":
			if m.showBack {
				ref := m.due[m.index]
				internal.Promote(&ref.Deck.Cards[ref.Idx])
				m.next()
			}

		case "n":
			if m.showBack {
				ref := m.due[m.index]
				internal.Reset(&ref.Deck.Cards[ref.Idx])
				m.next()
			}
		}
	}

	return m, nil
}

func (m *Model) next() {
	m.showBack = false
	m.index++

	if m.index >= len(m.due) {
		m.done = true
	}
}

func (m *Model) View() string {

	if m.done {
		return "\nAll done.\n"
	}

	ref := m.due[m.index]
	card := ref.Deck.Cards[ref.Idx]

	if !m.showBack {
		return fmt.Sprintf("\n\n%s\n", card.Front)
	}

	return fmt.Sprintf("\n\n%s\n\n%s\n", card.Front, card.Back)
}
