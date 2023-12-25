package size_test

import (
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sqrtOfMaxInt = 3037000500

func TestSquareDoesNotAllowNegativeValues(t *testing.T) {
	testCases := []struct {
		name string
		side int
	}{
		{
			name: "side cannot be negative",
			side: -10,
		},
		{
			name: "side cannot be zero",
			side: 0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := size.Square(testCase.side)
			assert.Error(t, err)
		})
	}
}

func TestSquareAllowsRandomPositiveValuesWithingLimits(t *testing.T) {
	testSize, err := size.Square(rand.Intn(sqrtOfMaxInt))
	assert.Nil(t, err)
	assert.NotNil(t, testSize)
}

func TestDefaultsToSqrtOfMaxIntWhenGivenALargerValue(t *testing.T) {
	testSize, err := size.Square(sqrtOfMaxInt + 1)
	assert.Nil(t, err)
	assert.Equal(t, sqrtOfMaxInt, testSize.Width())
	assert.Equal(t, sqrtOfMaxInt, testSize.Height())
}

func TestGetWidthAndHeight(t *testing.T) {
	dimension := 5
	testSize, err := size.Square(dimension)
	assert.Nil(t, err)

	assert.Equal(t, dimension, testSize.Width())
	assert.Equal(t, dimension, testSize.Height())
}

func TestCalculatesArea(t *testing.T) {
	dimension := 5
	testSize, err := size.Square(dimension)
	assert.Nil(t, err)

	assert.Equal(t, dimension*dimension, testSize.Area())
}
