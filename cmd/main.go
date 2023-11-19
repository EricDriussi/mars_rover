package cmd

import (
	"fmt"
	"mars_rover/internal/use_case/create"
	"mars_rover/internal/use_case/move"
)

// TODO: LIST OF THINGS!
// Persistence - WIP
// API
// GUI
func Sample1() {
	curiosity := create.Random()
	moveUseCase := move.For(curiosity)
	movementErrors := moveUseCase.MoveSequence("ffrfblf")

	for _, err := range movementErrors {
		fmt.Println(err)
	}

	fmt.Println("Rover finished moving")
}

func Sample2() {
	curiosity := create.Random()
	moveUseCase := move.For(curiosity)
	err := moveUseCase.MoveSequenceAborting("ffrfblf")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Rover finished moving")
	}
}
