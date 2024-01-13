package boundedRandomGameCreator

import (
	"errors"
	. "mars_rover/src/action/randomCreator"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
)

const (
	MinSize         = 8
	MaxSize         = 16
	MinObstacles    = 10
	MinBigObstacles = 5
)

type BoundedRandomCreator struct {
	repo            Repository
	minSize         int
	maxSize         int
	minObstacles    int
	minBigObstacles int
}

func With(repo Repository) *BoundedRandomCreator {
	return &BoundedRandomCreator{
		repo:            repo,
		minSize:         MinSize,
		maxSize:         MaxSize,
		minObstacles:    MinObstacles,
		minBigObstacles: MinBigObstacles,
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
		return CreatePlanet(RandomColor(), validSize, this.randomObstaclesWithin(validSize))
	})
}

func (this *BoundedRandomCreator) loopUntilValidSize() *Size {
	return LoopUntilNoError(func() (*Size, error) {
		randNumWithinLimits := rand.Intn(this.maxSize-this.minSize) + this.minSize
		return size.Square(randNumWithinLimits)
	})
}

func (this *BoundedRandomCreator) randomObstaclesWithin(size Size) Obstacles {
	list := obstacles.Empty()
	list = addBigObstacles(size, list, this.minBigObstacles)
	list = addRandomObstacles(size, list, this.calculateAmountOfRandomObstacles(size))

	return *list
}

func (this *BoundedRandomCreator) calculateAmountOfRandomObstacles(size Size) int {
	halfTheArea := size.Area() / 2
	betweenMinObstaclesAndHalfTheArea := rand.Intn(halfTheArea-this.minObstacles) + this.minObstacles
	amountOfRandomObstacles := betweenMinObstaclesAndHalfTheArea - this.minBigObstacles
	return amountOfRandomObstacles
}

func addBigObstacles(size Size, list *Obstacles, minAmount int) *Obstacles {
	for i := 0; i < minAmount; i++ {
		list = loopUntilAbleToAddBigObstacle(size, *list)
	}
	return list
}

func addRandomObstacles(size Size, list *Obstacles, minAmount int) *Obstacles {
	for i := 0; i < minAmount; i++ {
		list = LoopUntilAbleToAddRandomObstacle(size, *list)
	}
	return list
}

func loopUntilAbleToAddBigObstacle(size Size, list Obstacles) *Obstacles {
	err := errors.New("not added")
	for err != nil {
		err = list.Add(loopUntilValidBigObstacle(size))
	}
	return &list
}

func loopUntilValidBigObstacle(size Size) obstacle.Obstacle {
	return LoopUntilNoError(func() (obstacle.Obstacle, error) {
		return obstacle.CreateObstacle(*loopUntilValidCoordinates(size))
	})
}

func loopUntilValidCoordinates(size Size) *coordinates.Coordinates {
	return LoopUntilNoError(func() (*coordinates.Coordinates, error) {
		return randomCoordinatesForBigObstacle(size, obstacle.MaxSize())
	})
}

func randomCoordinatesForBigObstacle(size Size, maxObstacleSize int) (*coordinates.Coordinates, error) {
	betweenTwoAndMaxObstacleSize := rand.Intn(maxObstacleSize-2) + 2
	coords, err := coordinates.New(RandomCoordinateWithin(size))
	if err != nil {
		return nil, err
	}
	for i := 0; i < betweenTwoAndMaxObstacleSize; i++ {
		err := coords.Add(RandomCoordinateWithin(size))
		if err != nil {
			return nil, err
		}
	}
	return coords, nil
}
