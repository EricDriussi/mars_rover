package planet_test

import (
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateIfNoObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)
	_, err := planet.Create(*sizeLimit, obstaclesWithinBounds)

	assert.Nil(t, err)
}

func TestCannotCreateIfOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	_, err := planet.Create(*sizeLimit, oneObstacleOutOfBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfMoreThanOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.From(5, 5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	twoObstacleOutOfBounds := append(oneObstacleOutOfBounds, randomObstacleOutOf(*sizeLimit))
	_, err := planet.Create(*sizeLimit, twoObstacleOutOfBounds)

	assert.Error(t, err)
}

func generateThreeRandomObstaclesWithin(size size.Size) []obstacle.Obstacle {
	var obstacles []obstacle.Obstacle

	for i := 0; i < 3; i++ {
		randomObstacle := generateRandomObstacleWithin(size)
		obstacles = append(obstacles, randomObstacle)
	}

	return obstacles
}

func generateRandomObstacleWithin(size size.Size) obstacle.Obstacle {
	randomX := rand.Intn(size.Width)
	randomY := rand.Intn(size.Height)

	randomLocation, _ := location.From(randomX, randomY)
	return *obstacle.In(randomLocation)
}

func randomObstacleOutOf(size size.Size) obstacle.Obstacle {
	randomX := rand.Intn(99-size.Width+1) + size.Width
	randomY := rand.Intn(99-size.Height+1) + size.Height

	randomLocation, _ := location.From(randomX, randomY)
	return *obstacle.In(randomLocation)
}
