package direction_test

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	"mars_rover/internal/domain/location/direction"
	relativePosition "mars_rover/internal/domain/location/relative_position"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO.LM: Is this better than a file per direction and a test per function?

func TestNorth(t *testing.T) {
	north := &direction.North{}
	assert.Equal(t, north.CardinalPoint(), "N")
	assert.Equal(t, north.DirectionOnTheLeft().CardinalPoint(), "W")
	assert.Equal(t, north.DirectionOnTheRight().CardinalPoint(), "E")
	assert.Equal(t, north.RelativePositionAhead(), *relativePosition.New(coordinate2d.New(0, 1)))
	assert.Equal(t, north.RelativePositionBehind(), *relativePosition.New(coordinate2d.New(0, -1)))
}

func TestEast(t *testing.T) {
	east := &direction.East{}
	assert.Equal(t, east.CardinalPoint(), "E")
	assert.Equal(t, east.DirectionOnTheLeft().CardinalPoint(), "N")
	assert.Equal(t, east.DirectionOnTheRight().CardinalPoint(), "S")
	assert.Equal(t, east.RelativePositionAhead(), *relativePosition.New(coordinate2d.New(1, 0)))
	assert.Equal(t, east.RelativePositionBehind(), *relativePosition.New(coordinate2d.New(-1, 0)))
}

func TestSouth(t *testing.T) {
	south := &direction.South{}
	assert.Equal(t, south.CardinalPoint(), "S")
	assert.Equal(t, south.DirectionOnTheLeft().CardinalPoint(), "E")
	assert.Equal(t, south.DirectionOnTheRight().CardinalPoint(), "W")
	assert.Equal(t, south.RelativePositionAhead(), *relativePosition.New(coordinate2d.New(0, -1)))
	assert.Equal(t, south.RelativePositionBehind(), *relativePosition.New(coordinate2d.New(0, 1)))
}

func TestWest(t *testing.T) {
	west := &direction.West{}
	assert.Equal(t, west.CardinalPoint(), "W")
	assert.Equal(t, west.DirectionOnTheLeft().CardinalPoint(), "S")
	assert.Equal(t, west.DirectionOnTheRight().CardinalPoint(), "N")
	assert.Equal(t, west.RelativePositionAhead(), *relativePosition.New(coordinate2d.New(-1, 0)))
	assert.Equal(t, west.RelativePositionBehind(), *relativePosition.New(coordinate2d.New(1, 0)))
}
