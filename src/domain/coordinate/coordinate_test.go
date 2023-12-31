package coordinate_test

import (
	"mars_rover/src/domain/coordinate"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/coordinate/relativeCoordinate"
	. "mars_rover/src/domain/coordinate/relativeCoordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumsRelative(t *testing.T) {
	baseCoordinate := absoluteCoordinate.Build(1, 1)
	testCases := []struct {
		name     string
		relative *RelativeCoordinate
		expected *AbsoluteCoordinate
	}{
		{
			name:     "relative up",
			relative: relativeCoordinate.Up(),
			expected: absoluteCoordinate.Build(1, 2),
		},
		{
			name:     "relative right",
			relative: relativeCoordinate.Right(),
			expected: absoluteCoordinate.Build(2, 1),
		},
		{
			name:     "relative down",
			relative: relativeCoordinate.Down(),
			expected: absoluteCoordinate.Build(1, 0),
		},
		{
			name:     "relative left",
			relative: relativeCoordinate.Left(),
			expected: absoluteCoordinate.Build(0, 1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, coordinate.SumOf(*baseCoordinate, *testCase.relative), testCase.expected)
		})
	}
}
