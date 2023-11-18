package create

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/location"
	. "mars_rover/internal/domain/location/direction"
	. "mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover"
	"mars_rover/internal/domain/size"
	. "mars_rover/internal/domain/size"
	"math/rand"
)

// TODO.LM: Return error instead of panic?
func Random() Rover {
	randNum := rand.Intn(99) + 2
	randSize, err := size.Square(randNum)
	if err != nil {
		panic("Something went wrong :(")
	}

	randPlanet, err := rockyPlanet.Create(randomColor(), *randSize, randomObstaclesWithin(*randSize))
	if err != nil {
		panic("Something went wrong :(")
	}

	var randRover Rover
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

func randomObstaclesWithin(size Size) []Obstacle {
	var obstacles []Obstacle
	halfTheArea := size.Width() * size.Height() / 2
	betweenZeroAndHalfTheArea := rand.Intn(halfTheArea)
	for i := 0; i < betweenZeroAndHalfTheArea; i++ {
		obstacles = append(obstacles, rock.In(randomCoordinateWithin(size)))
	}
	return obstacles
}

func randomLocationWithin(size Size) (*location.Location, error) {
	return location.From(randomCoordinateWithin(size), randomDirection())
}

func randomCoordinateWithin(size Size) AbsoluteCoordinate {
	return *absoluteCoordinate.From(rand.Intn(size.Width()), rand.Intn(size.Height()))
}

func randomDirection() Direction {
	directions := []Direction{
		North{},
		East{},
		South{},
		West{},
	}
	return directions[rand.Intn(len(directions))]
}

func randomColor() string {
	colors := []string{
		"red",
		"blue",
		"green",
		"yellow",
		"black",
		"white",
		"pink",
		"purple",
		"orange",
		"brown",
	}
	return colors[rand.Intn(len(colors))]
}
