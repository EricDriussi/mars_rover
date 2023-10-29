package location_test

import (
	"mars_rover/internal/domain/location"
	relativePosition "mars_rover/internal/domain/relative_position"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportsFuturePositionWhenSizeLimitIsNotInvolved(t *testing.T) {
	testSize, _ := size.From(3, 3)
	testCases := []struct {
		name             string
		relativePosition *relativePosition.RelativePosition
		expectedX        int
		expectedY        int
	}{
		{
			name:             "increase X",
			relativePosition: relativePosition.New(1, 0),
			expectedX:        2,
			expectedY:        1,
		},
		{
			name:             "increase Y",
			relativePosition: relativePosition.New(0, 1),
			expectedX:        1,
			expectedY:        2,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			givenLocation, _ := location.From(1, 1)

			futureLocation := givenLocation.WillBeAt(*testCase.relativePosition, *testSize)

			expectedLocation, _ := location.From(testCase.expectedX, testCase.expectedY)
			assert.True(t, futureLocation.Equals(*expectedLocation))
		})
	}
}

func TestWrapsOn_Y(t *testing.T) {
	testSize, _ := size.From(3, 3)
	testCases := []struct {
		name             string
		relativePosition *relativePosition.RelativePosition
		startingY        int
		expectedY        int
	}{
		{
			name:             "Y over size",
			relativePosition: relativePosition.New(0, 1),
			startingY:        3,
			expectedY:        0,
		},
		{
			name:             "Y under size",
			relativePosition: relativePosition.New(0, -1),
			startingY:        0,
			expectedY:        3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			givenLocation, _ := location.From(1, testCase.startingY)

			futureLocation := givenLocation.WillBeAt(*testCase.relativePosition, *testSize)

			expectedLocation, _ := location.From(1, testCase.expectedY)
			assert.True(t, futureLocation.Equals(*expectedLocation))
		})
	}
}

func TestWrapsOn_X(t *testing.T) {
	testSize, _ := size.From(3, 3)
	testCases := []struct {
		name             string
		relativePosition *relativePosition.RelativePosition
		startingX        int
		expectedX        int
	}{
		{
			name:             "Y over size",
			relativePosition: relativePosition.New(1, 0),
			startingX:        3,
			expectedX:        0,
		},
		{
			name:             "Y under size",
			relativePosition: relativePosition.New(-1, 0),
			startingX:        0,
			expectedX:        3,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			givenLocation, _ := location.From(testCase.startingX, 1)

			futureLocation := givenLocation.WillBeAt(*testCase.relativePosition, *testSize)

			expectedLocation, _ := location.From(testCase.expectedX, 1)
			assert.True(t, futureLocation.Equals(*expectedLocation))
		})
	}
}
