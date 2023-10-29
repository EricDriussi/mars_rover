package cmd

import (
	"fmt"
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/service"
)

// TODO: check for var names left behind after the renames
// TODO: make sure tests are clear (limit of size is 0 not 1)
// TODO: review tests naming
// TODO: LIST OF THINGS!
// Collision detection
// Persistency
// API
// GUI
func Sample() {
	locationOfFirstObstacle, err := coordinate.From(3, 3)
	if err != nil {
		fmt.Println("Error creating coordinate:", err)
		return
	}

	locationOfSecondObstacle, err := coordinate.From(7, 7)
	if err != nil {
		fmt.Println("Error creating coordinate:", err)
		return
	}

	marsSize, err := size.From(10, 10)
	if err != nil {
		fmt.Println("Error creating size:", err)
		return
	}

	mars, err := planet.Create(*marsSize, []obstacle.Obstacle{*obstacle.In(locationOfFirstObstacle), *obstacle.In(locationOfSecondObstacle)})
	if err != nil {
		fmt.Println("Error creating planet:", err)
		return
	}

	landinglocation, err := coordinate.From(0, 0)
	if err != nil {
		fmt.Println("Error creating coordinate:", err)
		return
	}

	facingNorth := direction.North{}
	curiosity := rover.Land(*landinglocation, facingNorth, *mars)

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

		if curiosity.CheckObstacle() {
			fmt.Println("Obstacle detected!")
			break
		}
		fmt.Println("Rover completed command ", cmd)
	}

	fmt.Println("Rover completed all commands without issues!")
}
