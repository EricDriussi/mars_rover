package boundedRandomCreator

import (
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
)

const (
	MinSize      = 4
	MaxSize      = 20
	MinObstacles = 3
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
		minSize:      MinSize,
		maxSize:      MaxSize,
		minObstacles: MinObstacles,
	}
}

func (this *BoundedRandomCreator) Create() (Rover, *CreationError) {
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

func (this *BoundedRandomCreator) loopUntilPlanetCreated() Planet {
	return LoopUntilNoError(func() (Planet, error) {
		validSize := *this.loopUntilValidSize()
		return CreatePlanet(RandomColor(), validSize, *obstacles.FromList(this.randomObstaclesWithin(validSize)))
	})
}

func (this *BoundedRandomCreator) loopUntilValidSize() *Size {
	return LoopUntilNoError(func() (*Size, error) {
		randNumWithinLimits := rand.Intn(this.maxSize-this.minSize) + this.minSize
		return size.Square(randNumWithinLimits)
	})
}

func (this *BoundedRandomCreator) randomObstaclesWithin(size Size) []Obstacle {
	var list []Obstacle
	halfTheArea := size.Area() / 2
	betweenMinObstaclesAndHalfTheArea := rand.Intn(halfTheArea-this.minObstacles) + this.minObstacles
	for i := 0; i < betweenMinObstaclesAndHalfTheArea; i++ {
		list = append(list, LoopUntilValidObstacle(size))
	}
	return list
}
