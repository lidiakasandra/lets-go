package main

type Game struct {
	scenes      []scene
	items       []item
	activeScene int
}

func initiateGame() (Game, error) {
	scenes, err := loadScenes()
	if err != nil {
		return Game{}, err
	}
	return Game{
		scenes:      scenes,
		items:       []item{},
		activeScene: 0,
	}, nil
}
