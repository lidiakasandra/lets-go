package main

import (
	gloss "charm.land/lipgloss/v2"
)

type styles struct {
	header     gloss.Style
	footer     gloss.Style
	body       gloss.Style
	container  gloss.Style
	selected   gloss.Style
	unselected gloss.Style
}

func generateStyles(c UIconfiguration) styles {
	return styles{
		header: gloss.NewStyle().
			Padding(c.HeaderPaddingTop, c.HeaderPaddingRight, c.HeaderPaddingBottom, c.HeaderPaddingLeft).
			Bold(true),

		footer: gloss.NewStyle().
			Padding(c.FooterPaddingTop, c.FooterPaddingRight, c.FooterPaddingBottom, c.FooterPaddingLeft).
			Foreground(gloss.Color(c.FooterForegroundColor)),

		body: gloss.NewStyle(),

		container: gloss.NewStyle().
			Border(gloss.RoundedBorder()).
			BorderForeground(gloss.Color(c.BorderColor)).
			Padding(0, 0, 0, c.ContainerPaddingLeft),

		selected: gloss.NewStyle().Bold(true).
			Background(gloss.Color(c.SelectedBgColor)),

		unselected: gloss.NewStyle(),
	}
}
