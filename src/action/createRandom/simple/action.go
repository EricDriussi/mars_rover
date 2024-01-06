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
	repo Repository
}

func With(repo Repository) *SimpleRandomCreator {
	return &SimpleRandomCreator{
		repo: repo,
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
		return CreatePlanet(RandomColor(), validSize, this.randomObstaclesWithin(validSize))
	})
}

func (this *SimpleRandomCreator) loopUntilValidSize() *Size {
	return LoopUntilNoError(func() (*Size, error) {
		return size.Square(rand.Intn(420) + 1) // 420 has no meaning: could be left unbound, but tests would take too long ¯\_(ツ)_/¯
	})
}

func (this *SimpleRandomCreator) randomObstaclesWithin(size Size) Obstacles {
	list := obstacles.FromList()
	amountOfObstacles := rand.Intn(size.Area()) - 1 // leave at least a blank space for the rover
	for i := 0; i < amountOfObstacles; i++ {
		list.Add(LoopUntilValidObstacle(size))
	}
	return *list
}
