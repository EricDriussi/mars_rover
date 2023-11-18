package coordinate_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	"mars_rover/internal/domain/coordinate/relativeCoordinate"
	. "mars_rover/internal/domain/coordinate/relativeCoordinate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumsRelative(t *testing.T) {
	absCoord := absoluteCoordinate.From(1, 1)
	testCases := []struct {
		name               string
		relativeCoordinate *RelativeCoordinate
		expected           *AbsoluteCoordinate
	}{
		{
			name:               "relative up",
			relativeCoordinate: relativeCoordinate.New(0, 1),
			expected:           absoluteCoordinate.From(1, 2),
		},
		{
			name:               "relative right",
			relativeCoordinate: relativeCoordinate.New(1, 0),
			expected:           absoluteCoordinate.From(2, 1),
		},
		{
			name:               "relative down",
			relativeCoordinate: relativeCoordinate.New(0, -1),
			expected:           absoluteCoordinate.From(1, 0),
		},
		{
			name:               "relative left",
			relativeCoordinate: relativeCoordinate.New(-1, 0),
			expected:           absoluteCoordinate.From(0, 1),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			resultOfSum := coordinate.SumOf(*absCoord, *testCase.relativeCoordinate)
			assert.Equal(t, resultOfSum, testCase.expected)
		})
	}
}
