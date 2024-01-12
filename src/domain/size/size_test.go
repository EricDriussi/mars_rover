package size_test

import (
	"mars_rover/src/domain/size"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

const sqrtOfMaxInt = 3037000500

func TestDoesNotAllowValuesLessThanOne(t *testing.T) {
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
			testSize, err := size.Square(testCase.side)

			assert.Error(t, err)
			assert.Nil(t, testSize)
		})
	}
}

func TestAllowsRandomPositiveValuesWithinLimits(t *testing.T) {
	side := rand.Intn(sqrtOfMaxInt)

	testSize, err := size.Square(side)

	assert.Nil(t, err)
	assert.Equal(t, side, testSize.Width())
	assert.Equal(t, side, testSize.Height())
}

func TestDefaultsToSqrtOfMaxIntWhenGivenALargerValue(t *testing.T) {
	testSize, err := size.Square(sqrtOfMaxInt + 1)

	assert.Nil(t, err)
	assert.Equal(t, sqrtOfMaxInt, testSize.Width())
	assert.Equal(t, sqrtOfMaxInt, testSize.Height())
}

func TestCalculatesArea(t *testing.T) {
	side := 5

	testSize, err := size.Square(side)

	assert.Nil(t, err)
	assert.Equal(t, side*side, testSize.Area())
}
