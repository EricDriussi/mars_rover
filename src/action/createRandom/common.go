package randomCreator

import (
	"errors"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/planet"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/rover/wrappingCollidingRover"
	. "mars_rover/src/domain/size"
	"math/rand"
)

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
		return wrappingCollidingRover.LandFacing(uuid.New(), RandomDirection(), RandomCoordinateWithin(planet.Size()), planet)
	})
}

func RandomCoordinateWithin(size Size) AbsoluteCoordinate {
	return *absoluteCoordinate.Build(rand.Intn(size.Width()), rand.Intn(size.Height()))
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