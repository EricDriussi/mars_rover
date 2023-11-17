package cmd

import (
	"fmt"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/small_rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"mars_rover/internal/service/move"
)

// TODO: LIST OF THINGS!
// Persistency
// API
// GUI
func Sample() {
	marsSize, err := size.Square(10)
	if err != nil {
		fmt.Println("Error creating size:", err)
		return
	}

	mars, err := rockyPlanet.Create(*marsSize, []obstacle.Obstacle{rock.In(*absoluteCoordinate.From(3, 3)), rock.In(*absoluteCoordinate.From(7, 7))})
	if err != nil {
		fmt.Println("Error creating planet:", err)
		return
	}

	facingNorth := direction.North{}
	landinglocation, err := location.From(*absoluteCoordinate.From(0, 0), facingNorth)
	if err != nil {
		fmt.Println("Error creating location:", err)
		return
	}

	curiosity, err := rover.Land(*landinglocation, mars)
	if err != nil {
		fmt.Println("Could not land on selected location:", err)
		return
	}

	moveService := service.For(curiosity)
	movementErrors := moveService.MoveSequence("ffrfblf")

	for _, err := range movementErrors {
		fmt.Println(err)
	}

	fmt.Println("Rover finished moving")
}
