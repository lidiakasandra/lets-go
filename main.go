package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"gopkg.in/ini.v1"
)

type UIconfiguration struct {
	HeaderPaddingTop      int    `ini:"HeaderPaddingTop"`
	HeaderPaddingLeft     int    `ini:"HeaderPaddingLeft"`
	HeaderPaddingRight    int    `ini:"HeaderPaddingRight"`
	HeaderPaddingBottom   int    `ini:"HeaderPaddingBottom"`
	FooterPaddingTop      int    `ini:"FooterPaddingTop"`
	FooterPaddingLeft     int    `ini:"FooterPaddingLeft"`
	FooterPaddingRight    int    `ini:"FooterPaddingRight"`
	FooterPaddingBottom   int    `ini:"FooterPaddingBottom"`
	FooterText            string `ini:"FooterText"`
	ContainerPaddingLeft  int    `ini:"ContainerPaddingLeft"`
	BorderColor           string `ini:"BorderColor"`
	FooterForegroundColor string `ini:"FooterForegroundColor"`
	SelectedBgColor       string `ini:"SelectedBgColor"`
}

func loadConfig(path string) (config UIconfiguration, err error) {
	cfg, err := ini.LoadSources(
		ini.LoadOptions{
			IgnoreInlineComment: true,
		},
		path,
	)
	if err != nil {
		return config, err
	}
	if err := cfg.MapTo(&config); err != nil {
		return config, err
	}
	return config, nil
}

func main() {
	config, err := loadConfig("ui.conf")
	if err != nil {
		fmt.Printf("There's been an error loading config: %v", err)
		os.Exit(1)
	}
	p := tea.NewProgram(initialModel(config))
	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error in the app: %v", err)
		os.Exit(1)
	}
}
