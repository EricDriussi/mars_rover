package planetMap_test

import (
	"fmt"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet_map"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsCollision(t *testing.T) {
	x := 3
	y := 3
	planet := createPlanetWithObstacleIn(x, y)
	planetMap := planetMap.Of(*planet)
	obstaclePosition, _ := location.From(x, y)

	didCollide := planetMap.CheckCollision(*obstaclePosition)

	assert.True(t, didCollide)
}

func TestReportsNOCollision(t *testing.T) {
	planetSize, _ := size.From(5, 5)

	for x := 0; x <= planetSize.Width; x++ {
		for y := 0; y <= planetSize.Height; y++ {
			name := fmt.Sprintf("no collision in %d, %d", x, y)

			t.Run(name, func(t *testing.T) {
				testPosition, _ := location.From(x, y)
				testPlanet := createPlanetWithRandomObstaclesNotIn(*planetSize, *testPosition)

				planetMap := planetMap.Of(*testPlanet)
				didCollide := planetMap.CheckCollision(*testPosition)

				assert.False(t, didCollide)
			})
		}
	}
}

func createPlanetWithObstacleIn(x, y int) *planet.Planet {
	planetSize, _ := size.From(x+2, y+2)
	obstaclePosition, _ := location.From(x, y)
	planetObstacle := obstacle.In(obstaclePosition)
	planet, _ := planet.Create(*planetSize, []obstacle.Obstacle{*planetObstacle})
	return planet
}

func createPlanetWithRandomObstaclesNotIn(planetSize size.Size, exclude location.Location) *planet.Planet {
	maxNumOfObstacles := (planetSize.Height * planetSize.Width) - 1
	numObstacles := max(rand.Intn(maxNumOfObstacles), 1)

	var obstacles []obstacle.Obstacle

	for i := 0; i < numObstacles; i++ {
		randomPosition := getRandomPositionExcluding(planetSize, exclude)
		obstacles = append(obstacles, *obstacle.In(randomPosition))
	}

	planet, _ := planet.Create(planetSize, obstacles)
	return planet
}

func getRandomPositionExcluding(planetSize size.Size, exclude location.Location) *location.Location {
	for {
		randomPosition, _ := location.From(rand.Intn(planetSize.Width), rand.Intn(planetSize.Height))

		positionIsNotExcluded := !randomPosition.Equals(exclude)
		if positionIsNotExcluded {
			return randomPosition
		}
	}
}
