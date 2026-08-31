package main

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	gloss "charm.land/lipgloss/v2"
)

type model struct {
	game   *Game
	cursor int
	width  int
	height int
	config UIconfiguration
	styles styles
}

func initialModel(config UIconfiguration, styles styles, game *Game) model {
	return model{
		cursor: 0,
		game:   game,
		config: config,
		styles: styles,
	}
}
func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	currentScene := m.game.scenes[m.game.activeScene]
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < countSceneChoices(currentScene)-1 {
				m.cursor++
			}
		case "enter":
			if sceneHasChoices(currentScene) {
				transitionScene(m.game, currentScene.Choices[m.cursor].Transitions)
			} else {
				// if there are no choices, just go to next
				transitionScene(m.game, m.game.activeScene+1)
			}
		}
	}
	return m, nil
}
func (m model) View() tea.View {
	currentScene := m.game.scenes[m.game.activeScene]
	content := strings.Join(currentScene.Story, "\n") + "\n"
	// Print choices if they exist
	if sceneHasChoices(currentScene) {
		for i, choice := range currentScene.Choices {
			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
				content += m.styles.selected.Render(cursor+" "+choice.Text) + "\n"
			} else {
				content += m.styles.unselected.Render(cursor+" "+choice.Text) + "\n"
			}
		}
	}
	header := m.styles.header.Render(currentScene.Header)
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
