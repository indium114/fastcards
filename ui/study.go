package ui

import (
	"fmt"

	"github.com/stikypiston/fastcards/internal"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type keyMap struct {
	Flip key.Binding
	Yes  key.Binding
	No   key.Binding
	Quit key.Binding
}

func newKeyMap() keyMap {
	return keyMap{
		Flip: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "flip"),
		),
		Yes: key.NewBinding(
			key.WithKeys("y"),
			key.WithHelp("y", "correct"),
		),
		No: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "incorrect"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	}
}

type Model struct {
	due      []internal.DueRef
	index    int
	showBack bool
	done     bool

	keys keyMap
	help help.Model
}

func NewStudyModelFromRefs(refs []internal.DueRef) *Model {
	return &Model{
		due:  refs,
		keys: newKeyMap(),
		help: help.New(),
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

		switch {

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Flip):
			m.showBack = true

		case key.Matches(msg, m.keys.Yes):
			if m.showBack {
				ref := m.due[m.index]
				internal.Promote(&ref.Deck.Cards[ref.Idx])
				m.next()
			}

		case key.Matches(msg, m.keys.No):
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
		return "\nAll done.\nPress [q] to quit."
	}

	ref := m.due[m.index]
	card := ref.Deck.Cards[ref.Idx]

	content := card.Front
	if m.showBack {
		content = fmt.Sprintf("%s\n\n%s", card.Front, card.Back)
	}

	// Styled card box
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Render(content)

	progress := fmt.Sprintf("Card %d/%d\n\n", m.index+1, len(m.due))

	return progress + box + "\n\n" + m.help.View(m)
}

func (m *Model) ShortHelp() []key.Binding {
	return []key.Binding{
		m.keys.Flip,
		m.keys.Yes,
		m.keys.No,
		m.keys.Quit,
	}
}

func (m *Model) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			m.keys.Flip,
			m.keys.Yes,
			m.keys.No,
			m.keys.Quit,
		},
	}
}
