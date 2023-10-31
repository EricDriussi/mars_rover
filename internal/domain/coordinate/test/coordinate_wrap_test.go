package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TODO: should test WrapIfOutOf instead

func TestDoesNotWrapXIfWithinLimit(t *testing.T) {
	coord := coordinate.New(1, 1)
	coord.WrapXIfOutOf(2)
	assert.Equal(t, coord.X(), 1)
}

func TestWrapsXIfTooLarge(t *testing.T) {
	coord := coordinate.New(3, 1)
	coord.WrapXIfOutOf(2)
	assert.Equal(t, coord.X(), 0)
}

func TestWrapsXIfTooSmall(t *testing.T) {
	coord := coordinate.New(-1, 1)
	coord.WrapXIfOutOf(2)
	assert.Equal(t, coord.X(), 2)
}

func TestDoesNotWrapYIfWithinLimit(t *testing.T) {
	coord := coordinate.New(1, 1)
	coord.WrapYIfOutOf(2)
	assert.Equal(t, coord.Y(), 1)
}

func TestWrapsYIfTooLarge(t *testing.T) {
	coord := coordinate.New(1, 3)
	coord.WrapYIfOutOf(2)
	assert.Equal(t, coord.Y(), 0)
}

func TestWrapsYIfTooSmall(t *testing.T) {
	coord := coordinate.New(1, -1)
	coord.WrapYIfOutOf(2)
	assert.Equal(t, coord.Y(), 2)
}
