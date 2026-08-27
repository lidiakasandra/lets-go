package main

import (
	"encoding/json"
	"os"
)

type scene struct {
	Id      int      `json:"id"`
	Text    string   `json:"text"`
	Choices []choice `json:"choices"`
}
type item struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}
type choice struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Transitions int    `json:"transitions"`
	Consumes    []item `json:"consumes"`
	Obtains     []item `json:"obtains"`
}

func loadScenes() ([]scene, error) {
	data, err := os.ReadFile("assets/scenes.json")
	if err != nil {
		return nil, err
	}

	var scenes []scene
	if err := json.Unmarshal(data, &scenes); err != nil {
		return nil, err
	}
	return scenes, nil
}
