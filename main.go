package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	config, err := loadConfig("ui.conf")
	if err != nil {
		fmt.Printf("There's been an error loading config: %v", err)
		os.Exit(1)
	}
	styles := generateStyles(config)
	game, err := initiateGame()
	if err != nil {
		fmt.Printf("There's been an error initiating game: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(initialModel(config, styles, &game))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error in the app: %v", err)
		os.Exit(1)
	}
}
