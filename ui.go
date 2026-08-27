package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
)

type model struct {
	scenes      []scene
	activeScene int
	cursor      int
	width       int
	height      int
	config      UIconfiguration
	styles      styles
}

func initialModel(config UIconfiguration) model {
	scenes, err := loadScenes()
	styles := generateStyles(config)
	if err != nil {
		panic(err)
	}
	return model{
		cursor:      0,
		scenes:      scenes,
		activeScene: 0,
		config:      config,
		styles:      styles,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
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
			if m.activeScene < len(m.scenes)-1 {
				m.activeScene = m.activeScene + 1
				return m, nil
			}
		}
	}
	return m, nil
}
func (m model) View() tea.View {
	content := strings.Join(m.scenes[m.activeScene].Story, "\n") + "\n"
	// Print choices if they exist
	if len(m.scenes[m.activeScene].Choices) != 0 {
		for i, choice := range m.scenes[m.activeScene].Choices {
			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
				content += m.styles.selected.Render(cursor+choice.Text) + "\n"
			} else {
				content += m.styles.unselected.Render(cursor+choice.Text) + "\n"
			}
		}
	}

	header := m.styles.header.Render(m.scenes[m.activeScene].Header)
	footer := m.styles.footer.Render(m.config.FooterText)

	contentHeight := m.height - 2 - gloss.Height(header) - gloss.Height(footer)
	body := m.styles.body.Height(contentHeight).Render(content)

	ui := gloss.JoinVertical(
		gloss.Left,
		header, body, footer,
	)
	container := m.styles.container.Width(m.width).Height(m.height).Render(ui)
	v := tea.NewView(container)
	v.AltScreen = true

	return v
}
