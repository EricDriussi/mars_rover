package cmd

import (
	"fmt"
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/small_rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/service"
)

// TODO: LIST OF THINGS!
// Collision detection - DONE, but reporting collision is missing
// review tests - add interfaces and mocks? - Pending location and rover
// consider property based testing
// Persistency
// API
// GUI
func Sample() {
	marsSize, err := size.From(10, 10)
	if err != nil {
		fmt.Println("Error creating size:", err)
		return
	}

	mars, err := rockyPlanet.Create(*marsSize, []obstacle.Obstacle{rock.In(*coordinate.NewAbsolute(3, 3)), rock.In(*coordinate.NewAbsolute(7, 7))})
	if err != nil {
		fmt.Println("Error creating planet:", err)
		return
	}

	facingNorth := direction.North{}
	landinglocation, err := location.From(*coordinate.NewAbsolute(0, 0), facingNorth)
	if err != nil {
		fmt.Println("Error creating location:", err)
		return
	}

	curiosity, err := rover.Land(*landinglocation, mars)
	if err != nil {
		fmt.Println("Could not land on selected location:", err)
		return
	}

	commands := []string{"f", "f", "r", "f", "b", "l", "f"}

	moveService := service.For(curiosity)
	moveService.MoveSequence(commands)

	for _, cmd := range commands {
		if cmd == "f" {
			curiosity.MoveForward()
		} else if cmd == "b" {
			curiosity.MoveBackward()
		} else if cmd == "l" {
			curiosity.TurnLeft()
		} else if cmd == "r" {
			curiosity.TurnRight()
		}

		// if curiosity.CheckObstacle() {
		// 	fmt.Println("Obstacle detected!")
		// 	break
		// }
		fmt.Println("Rover completed command ", cmd)
	}

	fmt.Println("Rover completed all commands without issues!")
}
