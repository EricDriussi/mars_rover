package planetMap_test

import (
	"fmt"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/obstacle"
	rock "mars_rover/internal/domain/obstacle/smallRock"
	"mars_rover/internal/domain/planet"
	"mars_rover/internal/domain/planet/rockyPlanet"
	"mars_rover/internal/domain/rover/planetMap"
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsCollision(t *testing.T) {
	x := 3
	y := 3
	testPlanet := createPlanetWithObstacleIn(x, y)
	testMap := planetMap.Of(testPlanet)
	obstacleCoordinate := absoluteCoordinate.From(x, y)

	didCollide := testMap.CollidesWithObstacle(*obstacleCoordinate)

	assert.True(t, didCollide)
}

func TestReportsNOCollision(t *testing.T) {
	planetSize, _ := size.Square(5)

	for x := 0; x <= planetSize.Width(); x++ {
		for y := 0; y <= planetSize.Height(); y++ {
			name := fmt.Sprintf("no collision in %d, %d", x, y)

			t.Run(name, func(t *testing.T) {
				testCoordinate := absoluteCoordinate.From(x, y)
				testPlanet := createPlanetWithRandomObstaclesNotIn(*planetSize, *testCoordinate)

				testMap := planetMap.Of(testPlanet)
				didCollide := testMap.CollidesWithObstacle(*testCoordinate)

				assert.False(t, didCollide)
			})
		}
	}
}

func createPlanetWithObstacleIn(x, y int) planet.Planet {
	planetSize, _ := size.Square(y + 2)
	obstacleCoordinate := absoluteCoordinate.From(x, y)
	planetObstacle := rock.In(*obstacleCoordinate)
	testPlanet, _ := rockyPlanet.Create("testColor", *planetSize, []obstacle.Obstacle{planetObstacle})
	return testPlanet
}

func createPlanetWithRandomObstaclesNotIn(planetSize size.Size, exclude absoluteCoordinate.AbsoluteCoordinate) planet.Planet {
	maxNumOfObstacles := (planetSize.Height() * planetSize.Width()) - 1
	numObstacles := max(rand.Intn(maxNumOfObstacles), 1)

	var obstacles []obstacle.Obstacle

	for i := 0; i < numObstacles; i++ {
		randomCoordinate := getRandomCoordinateExcluding(planetSize, exclude)
		obstacles = append(obstacles, rock.In(randomCoordinate))
	}

	testPlanet, _ := rockyPlanet.Create("testColor", planetSize, obstacles)
	return testPlanet
}

func getRandomCoordinateExcluding(planetSize size.Size, exclude absoluteCoordinate.AbsoluteCoordinate) absoluteCoordinate.AbsoluteCoordinate {
	for {
		randomCoordinate := absoluteCoordinate.From(rand.Intn(planetSize.Width()), rand.Intn(planetSize.Height()))

		coordinateIsNotExcluded := !randomCoordinate.Equals(&exclude)
		if coordinateIsNotExcluded {
			return *randomCoordinate
		}
	}
}
