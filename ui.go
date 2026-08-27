package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	scenes      []scene
	activeScene int
	cursor      int
}

func initialModel() model {
	scenes, err := loadScenes()
	if err != nil {
		panic(err)
	}
	return model{
		cursor:      0,
		scenes:      scenes,
		activeScene: 0,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyPressMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.scenes[m.activeScene].Choices)-1 {
				m.cursor++
			}
		case "enter":
			m.activeScene = m.activeScene + 1
			return m, nil
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	return m, nil
}
func (m model) View() tea.View {
	// The header
	s := "\n"
	s += m.scenes[m.activeScene].Text + "\n\n"
	// Iterate over choices
	if len(m.scenes[m.activeScene].Choices) != 0 {
		for i, choice := range m.scenes[m.activeScene].Choices {
			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}
			s += fmt.Sprintf("%s [%s]\n", cursor, choice.Text)
		}
		// The footer
		s += fmt.Sprintf("\nHere is your current scene: [%d], and choice: %s\n", m.activeScene, m.scenes[m.activeScene].Choices[m.cursor].Text)
	} else {
		s += fmt.Sprintf("\nHere is your current scene: [%d]\n", m.activeScene)
	}
	s += "\nPress q to quit.\n"
	return tea.NewView(s)
}
