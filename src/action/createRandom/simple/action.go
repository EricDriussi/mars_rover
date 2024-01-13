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
	allObstacles := obstacles.Empty()
	for i := 0; i < rand.Intn(size.Width()); i++ {
		allObstacles = LoopUntilAbleToAddRandomObstacle(size, *allObstacles)
	}

	// This ensures the Rover always has a place to land
	// There is definitely a better way to do it
	allObstaclesMinusOne := obstacles.Empty()
	for i := 0; i < allObstacles.Amount()-1; i++ {
		err := allObstaclesMinusOne.Add(allObstacles.List()[i])
		if err != nil {
			panic(err)
		}
	}

	return *allObstaclesMinusOne
}
