package location_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/location"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAllowNegativeValues(t *testing.T) {
	testCases := []struct {
		name       string
		coordinate coordinate.Coordinate
	}{
		{
			name:       "neither X nor Y can be negative",
			coordinate: *coordinate.New(-1, -1),
		},
		{
			name:       "x cannot be negative",
			coordinate: *coordinate.New(-1, 1),
		},
		{
			name:       "y cannot be negative",
			coordinate: *coordinate.New(1, -1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := location.From(testCase.coordinate)
			assert.Error(t, err)
		})
	}
}

func TestEqualsBasedOnCoordinates(t *testing.T) {
	aLocation, _ := location.From(*coordinate.New(1, 1))
	anEqualLocation, _ := location.From(*coordinate.New(1, 1))

	areTheSame := aLocation.Equals(*anEqualLocation)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnCoordinates(t *testing.T) {
	aLocation, _ := location.From(*coordinate.New(1, 2))
	anEqualLocation, _ := location.From(*coordinate.New(2, 1))

	areTheSame := aLocation.Equals(*anEqualLocation)
	assert.False(t, areTheSame)
}
