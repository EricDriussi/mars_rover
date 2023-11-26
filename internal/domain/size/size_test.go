package size_test

import (
	"mars_rover/internal/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestSquareAllowsRandomPositiveValues(t *testing.T) {
	testSize, err := size.Square(rand.Int())
	assert.Nil(t, err)
	assert.NotNil(t, testSize)
}

func TestGetWidthAndHeight(t *testing.T) {
	dimension := 5
	testSize, err := size.Square(dimension)
	assert.Nil(t, err)

	assert.Equal(t, dimension, testSize.Width())
	assert.Equal(t, dimension, testSize.Height())
}
