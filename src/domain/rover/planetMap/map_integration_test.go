package planetMap_test

import (
	"fmt"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle"
	rock "mars_rover/src/domain/obstacle/smallRock"
	. "mars_rover/src/domain/planet"
	"mars_rover/src/domain/planet/planetWithObstacles"
	"mars_rover/src/domain/rover/planetMap"
	"mars_rover/src/domain/size"
	. "mars_rover/src/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsCollision(t *testing.T) {
	x := 3
	y := 3
	testPlanet := createPlanetWithObstacleIn(x, y)
	testMap := planetMap.OfPlanet(testPlanet)
	obstacleCoordinate := absoluteCoordinate.Build(x, y)

	didCollide := testMap.HasObstacleIn(*obstacleCoordinate)

	assert.True(t, didCollide)
}

func TestReportsNOCollision(t *testing.T) {
	planetSize, _ := size.Square(5)

	for x := 0; x <= planetSize.Width(); x++ {
		for y := 0; y <= planetSize.Height(); y++ {
			name := fmt.Sprintf("no collision in %d, %d", x, y)

			t.Run(name, func(t *testing.T) {
				testCoordinate := absoluteCoordinate.Build(x, y)
				testPlanet := createPlanetWithRandomObstaclesNotIn(*planetSize, *testCoordinate)

				testMap := planetMap.OfPlanet(testPlanet)
				didCollide := testMap.HasObstacleIn(*testCoordinate)

				assert.False(t, didCollide)
			})
		}
	}
}

func createPlanetWithObstacleIn(x, y int) Planet {
	planetSize, _ := size.Square(y + 2)
	obstacleCoordinate := absoluteCoordinate.Build(x, y)
	planetObstacle := rock.In(*obstacleCoordinate)
	testPlanet, _ := planetWithObstacles.Create("testColor", *planetSize, []Obstacle{&planetObstacle})
	return testPlanet
}

func createPlanetWithRandomObstaclesNotIn(planetSize Size, exclude AbsoluteCoordinate) Planet {
	maxNumOfObstacles := (planetSize.Height() * planetSize.Width()) - 1
	numObstacles := max(rand.Intn(maxNumOfObstacles), 1)

	var obstacles []Obstacle

	for i := 0; i < numObstacles; i++ {
		randomCoordinate := getRandomCoordinateExcluding(planetSize, exclude)
		smallRock := rock.In(randomCoordinate)
		obstacles = append(obstacles, &smallRock)
	}

	testPlanet, _ := planetWithObstacles.Create("testColor", planetSize, obstacles)
	return testPlanet
}

func getRandomCoordinateExcluding(planetSize Size, exclude AbsoluteCoordinate) AbsoluteCoordinate {
	for {
		randomCoordinate := absoluteCoordinate.Build(rand.Intn(planetSize.Width()), rand.Intn(planetSize.Height()))

		coordinateIsNotExcluded := !randomCoordinate.Equals(exclude)
		if coordinateIsNotExcluded {
			return *randomCoordinate
		}
	}
}
