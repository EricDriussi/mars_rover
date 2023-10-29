package location_test

import (
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAllowNegativeValues(t *testing.T) {
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "neither X nor Y can be negative",
			x:    -1,
			y:    -1,
		},
		{
			name: "x cannot be negative",
			x:    -1,
			y:    1,
		},
		{
			name: "y cannot be negative",
			x:    1,
			y:    -1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := location.From(testCase.x, testCase.y)
			assert.Error(t, err)
		})
	}
}

func TestEqualsBasedOnValues(t *testing.T) {
	aLocation, err := location.From(1, 1)
	assert.Nil(t, err)
	anEqualLocation, err := location.From(1, 1)
	assert.Nil(t, err)

	areTheSame := aLocation.Equals(*anEqualLocation)
	assert.True(t, areTheSame)
}

func TestNotEqualsBasedOnValues(t *testing.T) {
	aLocation, err := location.From(1, 1)
	assert.Nil(t, err)
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "neither x nor y can differ",
			x:    2,
			y:    2,
		},
		{
			name: "x cannot differ",
			x:    2,
			y:    1,
		},
		{
			name: "y cannot differ",
			x:    1,
			y:    2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			aDifferentLocation, _ := location.From(testCase.x, testCase.y)
			areTheSame := aLocation.Equals(*aDifferentLocation)
			assert.False(t, areTheSame)
		})
	}
}

func TestIsWithinAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(2, 2)
	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			validLocation, err := location.From(x, y)
			assert.Nil(t, err)

			assert.True(t, validLocation.IsWithin(*sizeLimit))
		}
	}
}

func TestIsNotWithinAGivenSize(t *testing.T) {
	sizeLimit, _ := size.From(3, 3)
	testCases := []struct {
		name string
		x    int
		y    int
	}{
		{
			name: "Both out of bounds",
			x:    4,
			y:    4,
		},
		{
			name: "X out of bounds",
			x:    4,
			y:    3,
		},
		{
			name: "Y out of bounds",
			x:    3,
			y:    4,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			outOfBoundsLocation, err := location.From(testCase.x, testCase.y)
			assert.Nil(t, err)

			assert.False(t, outOfBoundsLocation.IsWithin(*sizeLimit))
		})
	}
}
