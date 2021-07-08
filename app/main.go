package main

import (
	"github.com/austindoolittle/spacetraders/ui"
	"log"
)

func main() {
	engine, err := ui.NewUiEngine()
	if err != nil {
		log.Panicln(err)
	}

	err = engine.Run()
	if err != nil {
		log.Panicln(err)
	}
}
