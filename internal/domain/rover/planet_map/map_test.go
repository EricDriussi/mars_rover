package planetMap_test

import (
	"fmt"
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/rover/planet_map"
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
	obstacleCoordinate := coordinate.New(x, y)

	didCollide := planetMap.CheckCollision(obstacleCoordinate)

	assert.True(t, didCollide)
}

func TestReportsNOCollision(t *testing.T) {
	planetSize, _ := size.From(5, 5)

	for x := 0; x <= planetSize.Width; x++ {
		for y := 0; y <= planetSize.Height; y++ {
			name := fmt.Sprintf("no collision in %d, %d", x, y)

			t.Run(name, func(t *testing.T) {
				testCoordinate := coordinate.New(x, y)
				testPlanet := createPlanetWithRandomObstaclesNotIn(*planetSize, *testCoordinate)

				planetMap := planetMap.Of(*testPlanet)
				didCollide := planetMap.CheckCollision(testCoordinate)

				assert.False(t, didCollide)
			})
		}
	}
}

func createPlanetWithObstacleIn(x, y int) *planet.Planet {
	planetSize, _ := size.From(x+2, y+2)
	obstacleCoordinate := coordinate.New(x, y)
	planetObstacle := obstacle.In(obstacleCoordinate)
	planet, _ := planet.Create(*planetSize, []obstacle.Obstacle{planetObstacle})
	return planet
}

func createPlanetWithRandomObstaclesNotIn(planetSize size.Size, exclude coordinate.Coordinate2D) *planet.Planet {
	maxNumOfObstacles := (planetSize.Height * planetSize.Width) - 1
	numObstacles := max(rand.Intn(maxNumOfObstacles), 1)

	var obstacles []obstacle.Obstacle

	for i := 0; i < numObstacles; i++ {
		randomCoordinate := getRandomCoordinateExcluding(planetSize, exclude)
		obstacles = append(obstacles, obstacle.In(randomCoordinate))
	}

	planet, _ := planet.Create(planetSize, obstacles)
	return planet
}

func getRandomCoordinateExcluding(planetSize size.Size, exclude coordinate.Coordinate2D) *coordinate.Coordinate2D {
	for {
		randomCoordinate := coordinate.New(rand.Intn(planetSize.Width), rand.Intn(planetSize.Height))

		coordinateIsNotExcluded := !randomCoordinate.Equals(&exclude)
		if coordinateIsNotExcluded {
			return randomCoordinate
		}
	}
}
