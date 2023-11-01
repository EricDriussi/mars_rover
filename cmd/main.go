package cmd

import (
	"fmt"
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/obstacle/rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/service"
)

// TODO: LIST OF THINGS!
// Collision detection - DONE, but reporting collision is missing
// review tests - add interfaces and mocks?
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

	mars, err := rockyPlanet.Create(*marsSize, []obstacle.Obstacle{rock.In(coordinate2d.New(3, 3)), rock.In(coordinate2d.New(7, 7))})
	if err != nil {
		fmt.Println("Error creating planet:", err)
		return
	}

	facingNorth := direction.North{}
	landinglocation, err := location.From(coordinate2d.New(0, 0), facingNorth)
	if err != nil {
		fmt.Println("Error creating location:", err)
		return
	}

	curiosity := rover.Land(*landinglocation, mars)

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
