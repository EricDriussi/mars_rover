package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumsRelative(t *testing.T) {
	absoluteCoordinate := coordinate.NewAbsolute(1, 1)
	testCases := []struct {
		name               string
		relativeCoordinate *coordinate.RelativeCoordinate
		expected           *coordinate.AbsoluteCoordinate
	}{
		{
			name:               "relative up",
			relativeCoordinate: coordinate.RelativeFrom(0, 1),
			expected:           coordinate.NewAbsolute(1, 2),
		},
		{
			name:               "relative right",
			relativeCoordinate: coordinate.RelativeFrom(1, 0),
			expected:           coordinate.NewAbsolute(2, 1),
		},
		{
			name:               "relative down",
			relativeCoordinate: coordinate.RelativeFrom(0, -1),
			expected:           coordinate.NewAbsolute(1, 0),
		},
		{
			name:               "relative left",
			relativeCoordinate: coordinate.RelativeFrom(-1, 0),
			expected:           coordinate.NewAbsolute(0, 1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultOfSum := coordinate.SumOf(*absoluteCoordinate, *testCase.relativeCoordinate)
			assert.Equal(t, resultOfSum, testCase.expected)
		})
	}
}
