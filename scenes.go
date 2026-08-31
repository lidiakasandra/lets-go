package main

import (
	"encoding/json"
	"os"
)

type scene struct {
	Id         int      `json:"id"`
	Header     string   `json:"header"`
	Story      []string `json:"story"`
	Choices    []choice `json:"choices"`
	Transition int      `json:"transition"`
}
type choice struct {
	Id          int     `json:"id"`
	Text        string  `json:"text"`
	Transitions int     `json:"transitions"`
	Consumes    []item  `json:"consumes"`
	Obtains     []item  `json:"obtains"`
	Achievement string  `json:"achievement"`
	Skills      []skill `json:"skills"`
	Crumb       string  `json:"crumb"`
}
type skill struct {
	Id    string `json:"id"`
	Value int    `json:"value"`
}
type item struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
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

// If the scene that you want to transition to exists, perform transition
func transitionScene(game *Game, i choice, c *int) {
	if i.Transitions > 0 { // 0 is Id of home screen and nothing will be able to navigate to it
		if validScene(game.scenes, i.Transitions) {
			game.activeScene = i.Transitions
			*c = 0
			game.sceneCrumb = i.Crumb
		} else {
			// if Id is invalid, for now do nothing
		}
	} else { //
		// you are trying to access home screen, for now do nothing
	}
}

// From the array of scenes, check if any of them has given Id, basically check if scene of this Id exists
func validScene(scenes []scene, i int) bool {
	for _, scene := range scenes {
		if scene.Id == i {
			return true
		}
	}
	return false
}

// For particular scene, check if it has valid choices
func sceneHasChoices(scene scene) bool {
	if countSceneChoices(scene) > 0 {
		return true
	}
	return false
}

func countSceneChoices(scene scene) int {
	return len(scene.Choices)
}
