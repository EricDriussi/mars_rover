package create

import (
	"github.com/google/uuid"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
)

// TODO: should there be only one action?
type CreateAction struct {
	repo Repository
}

func For(repo Repository) *CreateAction {
	return &CreateAction{
		repo: repo,
	}
}

func (this *CreateAction) Random() (Rover, error) {
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
		randRover, err = wrappingCollidingRover.LandFacing(uuid.New(), randomDirection(), randomCoordinateWithin(*randSize), randPlanet)
		if err == nil {
			couldNotLand = false
		}
	}

	planetId, err := this.repo.AddPlanet(randPlanet)
	err = this.repo.AddRover(randRover, planetId)
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
