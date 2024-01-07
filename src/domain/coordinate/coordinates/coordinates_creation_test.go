package coordinates_test

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/coordinates"
	"testing"
)

func TestFiltersOutDuplicateCoordinates(t *testing.T) {
	coords, err := coordinates.New(
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 1),
		*absoluteCoordinate.Build(1, 3),
	)

	assert.Nil(t, err)
	assert.Len(t, coords.List(), 2)
}

func TestDoesNotCreateWithNoCoordinates(t *testing.T) {
	coords, err := coordinates.New()

	assert.Nil(t, coords)
	assert.Error(t, err)
}
