package cmd

import (
	"fmt"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/service/move"
)

// TODO: LIST OF THINGS!
// Persistence - WIP
// API
// GUI
func Sample() {
	marsSize, err := size.Square(10)
	if err != nil {
		fmt.Println("Error creating size:", err)
		return
	}

	mars, err := rockyPlanet.Create("red", *marsSize, []Obstacle{rock.In(*absoluteCoordinate.From(3, 3)), rock.In(*absoluteCoordinate.From(7, 7))})
	if err != nil {
		fmt.Println("Error creating planet:", err)
		return
	}

	coordinate := absoluteCoordinate.From(0, 0)

	curiosity, err := wrappingCollidingRover.Land(*coordinate, mars)
	if err != nil {
		fmt.Println("Could not land on selected spot:", err)
		return
	}

	moveService := service.For(curiosity)
	movementErrors := moveService.MoveSequence("ffrfblf")

	for _, err := range movementErrors {
		fmt.Println(err)
	}

	fmt.Println("Rover finished moving")
}
