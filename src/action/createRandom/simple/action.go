package simpleRandomCreator

import (
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
)

type SimpleRandomCreator struct {
	repo    Repository
	maxSize int
}

func With(repo Repository, maxSize int) *SimpleRandomCreator {
	return &SimpleRandomCreator{
		repo:    repo,
		maxSize: maxSize,
	}
}

func (this *SimpleRandomCreator) Create() (Rover, *CreationError) {
	randPlanet := this.loopUntilPlanetCreated()
	randRover := LoopUntilRoverLanded(randPlanet)

	planetId, addPlanetErr := this.repo.AddPlanet(randPlanet)
	if addPlanetErr != nil {
		return nil, GameNotPersistedErr(addPlanetErr)
	}
	addRoverErr := this.repo.AddRover(randRover, planetId)
	if addRoverErr != nil {
		return nil, GameNotPersistedErr(addPlanetErr)
	}
	return randRover, nil
}

func (this *SimpleRandomCreator) loopUntilPlanetCreated() Planet {
	return LoopUntilNoError(func() (Planet, error) {
		validSize := *this.loopUntilValidSize()
		return CreatePlanet(RandomColor(), validSize, randomObstaclesWithin(validSize))
	})
}

func (this *SimpleRandomCreator) loopUntilValidSize() *Size {
	return LoopUntilNoError(func() (*Size, error) {
		return size.Square(rand.Intn(this.maxSize))
	})
}

func randomObstaclesWithin(size Size) Obstacles {
	list := obstacles.Empty()
	amountOfObstacles := rand.Intn(size.Area()) - 1 // leave at least a blank space for the rover
	for i := 0; i < amountOfObstacles; i++ {
		list = LoopUntilAbleToAddRandomObstacle(size, *list)
	}
	return *list
}
