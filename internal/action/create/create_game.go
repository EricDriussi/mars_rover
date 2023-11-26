package create

import (
	. "mars_rover/internal/domain"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet/rockyPlanet"
	. "mars_rover/internal/domain/rover"
	. "mars_rover/internal/domain/rover/direction"
	"mars_rover/internal/domain/rover/wrappingCollidingRover"
	"mars_rover/internal/domain/size"
	. "mars_rover/internal/domain/size"
	"math/rand"
)

func Random(repository Repository) (Rover, error) {
	randNum := rand.Intn(19) + 4
	randSize, err := size.Square(randNum)
	if err != nil {
		return nil, err
	}

	randPlanet, err := rockyPlanet.Create(randomColor(), *randSize, randomObstaclesWithin(*randSize))
	if err != nil {
		return nil, err
	}

	var randRover Rover
	couldNotLand := true
	for couldNotLand {
		randRover, err = wrappingCollidingRover.LandFacing(randomDirection(), randomCoordinateWithin(*randSize), randPlanet)
		if err == nil {
			couldNotLand = false
		}
	}

	err = repository.SaveGame(randRover, randPlanet)
	if err != nil {
		return nil, err
	}
	return randRover, nil
}

func randomObstaclesWithin(size Size) []Obstacle {
	var obstacles []Obstacle
	halfTheArea := size.Width() * size.Height() / 2
	betweenZeroAndHalfTheArea := rand.Intn(halfTheArea)
	for i := 0; i < betweenZeroAndHalfTheArea; i++ {
		smallRock := rock.In(randomCoordinateWithin(size))
		obstacles = append(obstacles, &smallRock)
	}
	return obstacles
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
