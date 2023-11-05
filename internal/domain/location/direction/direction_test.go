package direction_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location/direction"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO.LM: Is this better than a file per direction and a test per function?

func TestNorth(t *testing.T) {
	north := &direction.North{}
	assert.Equal(t, north.Degree(), 90)
	assert.Equal(t, north.DirectionOnTheLeft().Degree(), 0)
	assert.Equal(t, north.DirectionOnTheRight().Degree(), 180)
	assert.Equal(t, north.RelativePositionAhead(), *coordinate.RelativeFrom(0, 1))
	assert.Equal(t, north.RelativePositionBehind(), *coordinate.RelativeFrom(0, -1))
}

func TestEast(t *testing.T) {
	east := &direction.East{}
	assert.Equal(t, east.Degree(), 180)
	assert.Equal(t, east.DirectionOnTheLeft().Degree(), 90)
	assert.Equal(t, east.DirectionOnTheRight().Degree(), 270)
	assert.Equal(t, east.RelativePositionAhead(), *coordinate.RelativeFrom(1, 0))
	assert.Equal(t, east.RelativePositionBehind(), *coordinate.RelativeFrom(-1, 0))
}

func TestSouth(t *testing.T) {
	south := &direction.South{}
	assert.Equal(t, south.Degree(), 270)
	assert.Equal(t, south.DirectionOnTheLeft().Degree(), 180)
	assert.Equal(t, south.DirectionOnTheRight().Degree(), 0)
	assert.Equal(t, south.RelativePositionAhead(), *coordinate.RelativeFrom(0, -1))
	assert.Equal(t, south.RelativePositionBehind(), *coordinate.RelativeFrom(0, 1))
}

func TestWest(t *testing.T) {
	west := &direction.West{}
	assert.Equal(t, west.Degree(), 0)
	assert.Equal(t, west.DirectionOnTheLeft().Degree(), 270)
	assert.Equal(t, west.DirectionOnTheRight().Degree(), 90)
	assert.Equal(t, west.RelativePositionAhead(), *coordinate.RelativeFrom(-1, 0))
	assert.Equal(t, west.RelativePositionBehind(), *coordinate.RelativeFrom(1, 0))
}
