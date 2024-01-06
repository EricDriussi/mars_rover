package randomCreator

import (
	"errors"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/obstacle"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/size"
	"math/rand"
)

// TODO: test these functions?
func RandomColor() string {
	colors := []string{
		"red",
		"blue",
		"green",
	}
	return colors[rand.Intn(len(colors))]
}

func LoopUntilRoverLanded(planet Planet) Rover {
	return LoopUntilNoError(func() (*WrappingCollidingRover, error) {
		return wrappingCollidingRover.LandFacing(uuid.New(), RandomDirection(), randomCoordinateWithin(planet.Size()), planet)
	})
}

func randomCoordinateWithin(size Size) AbsoluteCoordinate {
	return *absoluteCoordinate.Build(rand.Intn(size.Width()), rand.Intn(size.Height()))
}

func LoopUntilValidObstacle(size Size) Obstacle {
	return LoopUntilNoError(func() (Obstacle, error) {
		coords, err := coordinates.New(randomCoordinatesWithin(size, obstacle.MaxAmountOfCoords())...)
		if err != nil {
			panic(err) // TODO: find a better way
		}
		return obstacle.CreateObstacle(*coords)
	})
}

func randomCoordinatesWithin(size Size, maxObstacleSize int) []AbsoluteCoordinate {
	betweenOneAndMaxObstacleSize := rand.Intn(maxObstacleSize-1) + 1
	var coords []AbsoluteCoordinate
	for i := 0; i < betweenOneAndMaxObstacleSize; i++ {
		coords = append(coords, randomCoordinateWithin(size))
	}
	return coords
}

func RandomDirection() Direction {
	directions := []Direction{
		North{},
		East{},
		South{},
		West{},
	}
	return directions[rand.Intn(len(directions))]
}

func LoopUntilNoError[T interface{}](create func() (T, error)) T {
	var t T
	err := errors.New("not created")
	for err != nil {
		t, err = create()
	}
	return t
}
