package planetWithObstacles_test

import (
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	"mars_rover/src/domain/planet/planetWithObstacles"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateIfNoObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, obstaclesWithinBounds)

	assert.Nil(t, err)
}

func TestCannotCreateIfOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, oneObstacleOutOfBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfMoreThanOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	twoObstacleOutOfBounds := append(oneObstacleOutOfBounds, randomObstacleOutOf(*sizeLimit))
	_, err := planetWithObstacles.Create("testColor", *sizeLimit, twoObstacleOutOfBounds)

	assert.Error(t, err)
}

func generateThreeRandomObstaclesWithin(size Size) []Obstacle {
	var obstacles []Obstacle
	for i := 0; i < 3; i++ {
		randomObstacle := generateRandomObstacleWithin(size)
		obstacles = append(obstacles, randomObstacle)
	}
	return obstacles
}

func generateRandomObstacleWithin(size Size) Obstacle {
	randomX := rand.Intn(size.Width())
	randomY := rand.Intn(size.Height())
	smallRock := rock.In(*absoluteCoordinate.Build(randomX, randomY))
	return &smallRock
}

func randomObstacleOutOf(size Size) Obstacle {
	randomX := rand.Intn(99-size.Width()) + size.Width() + 1
	randomY := rand.Intn(99-size.Height()) + size.Height() + 1
	smallRock := rock.In(*absoluteCoordinate.Build(randomX, randomY))
	return &smallRock
}
