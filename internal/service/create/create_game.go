package create

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/location/direction"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/small_rock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	"math/rand"
)

// TODO.LM: Return error instead of panic?
func Random() rover.Rover {
	randNum := rand.Intn(99) + 2
	randSize, err := size.From(randNum, randNum)
	if err != nil {
		panic("Something went wrong :(")
	}

	randPlanet, err := rockyPlanet.Create(*randSize, randomObstaclesWithin(*randSize))
	if err != nil {
		panic("Something went wrong :(")
	}

	var randRover rover.Rover
	couldNotLand := true
	for couldNotLand {
		landinglocation, err := randomLocationWithin(*randSize)
		if err != nil {
			panic("Something went wrong :(")
		}

		randRover, err = rover.Land(*landinglocation, randPlanet)
		if err == nil {
			couldNotLand = false
		}
	}
	// TODO: Persist planet and rover
	return randRover
}

func randomObstaclesWithin(size size.Size) []obstacle.Obstacle {
	var obstacles []obstacle.Obstacle
	halfTheArea := size.Width * size.Height / 2
	betweenZeroAndHalfTheArea := rand.Intn(halfTheArea)
	for i := 0; i < betweenZeroAndHalfTheArea; i++ {
		obstacles = append(obstacles, rock.In(randomCoordinateWithin(size)))
	}
	return obstacles
}

func randomLocationWithin(size size.Size) (*location.Location, error) {
	return location.From(randomCoordinateWithin(size), randomDirection())
}

func randomCoordinateWithin(size size.Size) absoluteCoordinate.AbsoluteCoordinate {
	return *absoluteCoordinate.From(rand.Intn(size.Width), rand.Intn(size.Height))
}

func randomDirection() direction.Direction {
	directions := []direction.Direction{
		direction.North{},
		direction.East{},
		direction.South{},
		direction.West{},
	}
	return directions[rand.Intn(len(directions))]
}
