package direction_test

import (
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/rover/direction"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO.LM: not sure if this approach is better than a file per direction and a test per function ¯\_(ツ)_/¯

func TestNorth(t *testing.T) {
	north := &North{}
	assert.Equal(t, north.CardinalPoint(), "N")
	assert.Equal(t, north.DirectionOnTheLeft().CardinalPoint(), "W")
	assert.Equal(t, north.DirectionOnTheRight().CardinalPoint(), "E")
	assert.Equal(t, north.RelativeCoordinateAhead(), *relativeCoordinate.New(0, 1))
	assert.Equal(t, north.RelativeCoordinateBehind(), *relativeCoordinate.New(0, -1))
}

func TestEast(t *testing.T) {
	east := &East{}
	assert.Equal(t, east.CardinalPoint(), "E")
	assert.Equal(t, east.DirectionOnTheLeft().CardinalPoint(), "N")
	assert.Equal(t, east.DirectionOnTheRight().CardinalPoint(), "S")
	assert.Equal(t, east.RelativeCoordinateAhead(), *relativeCoordinate.New(1, 0))
	assert.Equal(t, east.RelativeCoordinateBehind(), *relativeCoordinate.New(-1, 0))
}

func TestSouth(t *testing.T) {
	south := &South{}
	assert.Equal(t, south.CardinalPoint(), "S")
	assert.Equal(t, south.DirectionOnTheLeft().CardinalPoint(), "E")
	assert.Equal(t, south.DirectionOnTheRight().CardinalPoint(), "W")
	assert.Equal(t, south.RelativeCoordinateAhead(), *relativeCoordinate.New(0, -1))
	assert.Equal(t, south.RelativeCoordinateBehind(), *relativeCoordinate.New(0, 1))
}

func TestWest(t *testing.T) {
	west := &West{}
	assert.Equal(t, west.CardinalPoint(), "W")
	assert.Equal(t, west.DirectionOnTheLeft().CardinalPoint(), "S")
	assert.Equal(t, west.DirectionOnTheRight().CardinalPoint(), "N")
	assert.Equal(t, west.RelativeCoordinateAhead(), *relativeCoordinate.New(-1, 0))
	assert.Equal(t, west.RelativeCoordinateBehind(), *relativeCoordinate.New(1, 0))
}
