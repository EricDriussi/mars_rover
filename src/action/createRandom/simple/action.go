package simpleRandomCreator

import (
	"errors"
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/planet/rockyPlanet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/rover/wrappingCollidingRover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"

	"github.com/google/uuid"
)

type BoundedRandomCreator struct {
	repo Repository
}

func With(repo Repository) *BoundedRandomCreator {
	return &BoundedRandomCreator{
		repo: repo,
	}
}

func (this *BoundedRandomCreator) Create() (Rover, *CreationError) {
	randPlanet := this.loopUntilPlanetCreated()
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

func (this *BoundedRandomCreator) loopUntilPlanetCreated() *RockyPlanet {
	return loopUntilNoError(func() (*RockyPlanet, error) {
		validSize := *this.loopUntilValidSize()
		return rockyPlanet.Create(randomColor(), validSize, this.randomObstaclesWithin(validSize))
	})
}

func (this *BoundedRandomCreator) loopUntilValidSize() *Size {
	return loopUntilNoError(func() (*Size, error) {
		return size.Square(rand.Intn(420)) // could be left unbound, but tests would take too long ¯\_(ツ)_/¯
	})
}

func (this *BoundedRandomCreator) randomObstaclesWithin(size Size) []Obstacle {
	var obstacles []Obstacle
	amountOfObstacles := rand.Intn(size.Area() - 1) // leave at least a blank space for the rover
	for i := 0; i < amountOfObstacles; i++ {
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
	}
	return colors[rand.Intn(len(colors))]
}

func (this *BoundedRandomCreator) loopUntilRoverLanded(planet Planet) Rover {
	return loopUntilNoError(func() (*WrappingCollidingRover, error) {
		return wrappingCollidingRover.LandFacing(uuid.New(), randomDirection(), randomCoordinateWithin(planet.Size()), planet)
	})
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

func loopUntilNoError[T *RockyPlanet | *Size | *WrappingCollidingRover](create func() (T, error)) T {
	var t T
	err := errors.New("not created")
	for err != nil {
		t, err = create()
	}
	return t
}
