package bounded_random_creator

import (
	"errors"
	"github.com/google/uuid"
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
)

type BoundedRandomCreator struct {
	repo         Repository
	minSize      int
	maxSize      int
	minObstacles int
}

func With(repo Repository) *BoundedRandomCreator {
	return &BoundedRandomCreator{
		repo:         repo,
		minSize:      4,
		maxSize:      20,
		minObstacles: 3,
	}
}

func (this *BoundedRandomCreator) Create() (Rover, *CreationError) {
	randSize := this.loopUntilValidSize()
	randPlanet := this.loopUntilPlanetCreated(*randSize)
	randRover := this.loopUntilRoverLanded(randPlanet)

	planetId, err := this.repo.AddPlanet(randPlanet)
	if err != nil {
		return nil, GameNotPersistedErr(err)
	}
	err = this.repo.AddRover(randRover, planetId)
	if err != nil {
		return nil, GameNotPersistedErr(err)
	}
	return randRover, nil
}

func (this *BoundedRandomCreator) loopUntilValidSize() *Size {
	var randSize *Size
	err := errors.New("not created yet")
	for err != nil {
		randNum := rand.Intn(this.maxSize-this.minSize) + this.minSize
		randSize, err = size.Square(randNum)
	}
	return randSize
}

func (this *BoundedRandomCreator) loopUntilPlanetCreated(size Size) Planet {
	var randPlanet Planet
	err := errors.New("not created yet")
	for err != nil {
		randPlanet, err = rockyPlanet.Create(randomColor(), size, this.randomObstaclesWithin(size))
	}
	return randPlanet
}

func (this *BoundedRandomCreator) randomObstaclesWithin(size Size) []Obstacle {
	var obstacles []Obstacle
	halfTheArea := size.Area() / 2
	betweenMinAndHalfTheArea := rand.Intn(halfTheArea-this.minObstacles) + this.minObstacles
	for i := 0; i < betweenMinAndHalfTheArea; i++ {
		smallRock := rock.In(randomCoordinateWithin(size))
		obstacles = append(obstacles, &smallRock)
	}
	return obstacles
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

func (this *BoundedRandomCreator) loopUntilRoverLanded(planet Planet) Rover {
	var randRover Rover
	err := errors.New("not created yet")
	for err != nil {
		randRover, err = wrappingCollidingRover.LandFacing(uuid.New(), randomDirection(), randomCoordinateWithin(planet.Size()), planet)
	}
	return randRover
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
