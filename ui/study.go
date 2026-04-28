package ui

import (
	"fmt"

	"github.com/indium114/fastcards/internal"

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

var cardsStudied int
var sessionMsg string

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
				m.next(true)
			}

		case key.Matches(msg, m.keys.No):
			if m.showBack {
				ref := m.due[m.index]
				internal.Reset(&ref.Deck.Cards[ref.Idx])
				m.next(false)
			}
		}
	}

	return m, nil
}

func (m *Model) next(correct bool) {
	m.showBack = false
	m.index++
	cardsStudied++

	sessionMsg = ""

	xp, _ := internal.LoadXP()

	if correct {
		if cardsStudied%5 == 0 {
			xp += 20
			sessionMsg = fmt.Sprintf(" +20 XP! Studied %d cards", cardsStudied)
		}
	} else {
		xp -= 5
		if xp < 0 {
			xp = 0
		}
		sessionMsg = "󱕤 -5 XP for incorrect answer"
	}

	internal.SaveXP(xp)

	if m.index >= len(m.due) && correct {
		xp += 100
		internal.SaveXP(xp)
		if sessionMsg != "" {
			sessionMsg += " "
		}
		sessionMsg += " +100 XP! Finished all due cards"
	}

	if m.index >= len(m.due) {
		m.done = true
	}
}

func (m *Model) View() string {

	if m.done {
		msg := "\nAll done.\nPress [q] to quit."
		if sessionMsg != "" {
			msg = sessionMsg + "\n" + msg
		}
		return msg
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

	msg := progress + box

	if sessionMsg != "" {
		msg = sessionMsg + "\n\n" + msg
	}

	return msg + "\n\n" + m.help.View(m)
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
