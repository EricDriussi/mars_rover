package rocky_planet_test

import (
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	rockyPlanet "mars_rover/internal/domain/planet/rocky_planet"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanCreateIfNoObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)
	_, err := rockyPlanet.Create(*sizeLimit, obstaclesWithinBounds)

	assert.Nil(t, err)
}

func TestCannotCreateIfOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	_, err := rockyPlanet.Create(*sizeLimit, oneObstacleOutOfBounds)

	assert.Error(t, err)
}

func TestCannotCreateIfMoreThanOneObstacleIsOutOfBounds(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	obstaclesWithinBounds := generateThreeRandomObstaclesWithin(*sizeLimit)

	oneObstacleOutOfBounds := append(obstaclesWithinBounds, randomObstacleOutOf(*sizeLimit))
	twoObstacleOutOfBounds := append(oneObstacleOutOfBounds, randomObstacleOutOf(*sizeLimit))
	_, err := rockyPlanet.Create(*sizeLimit, twoObstacleOutOfBounds)

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
	randomX := rand.Intn(size.Width())
	randomY := rand.Intn(size.Height())

	randomLocation := absoluteCoordinate.From(randomX, randomY)
	return rock.In(*randomLocation)
}

func randomObstacleOutOf(size size.Size) obstacle.Obstacle {
	randomX := rand.Intn(99-size.Width()) + size.Width() + 1
	randomY := rand.Intn(99-size.Height()) + size.Height() + 1

	randomLocation := absoluteCoordinate.From(randomX, randomY)
	return rock.In(*randomLocation)
}
