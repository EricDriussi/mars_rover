package small_rock_test

import (
	"mars_rover/internal/domain/coordinate/test"
	smallRock "mars_rover/internal/domain/obstacle/small_rock"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 2, Height: 2}
	mockCoordinate := new(test.MockCoordinate)

	for x := 0; x <= sizeLimit.Width; x++ {
		for y := 0; y <= sizeLimit.Height; y++ {
			mockCoordinate.On("X").Return(x)
			mockCoordinate.On("Y").Return(y)
			rock := smallRock.In(mockCoordinate)

			assert.False(t, rock.IsBeyond(*sizeLimit))
			mockCoordinate.AssertCalled(t, "X")
			mockCoordinate.AssertCalled(t, "Y")
		}
	}
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 3, Height: 3}
	mockCoordinate := new(test.MockCoordinate)
	rock := smallRock.In(mockCoordinate)
	testCases := []struct {
		name string
		x    *mock.Call
		y    *mock.Call
	}{
		{
			name: "both out of bounds",
			x:    mockCoordinate.On("X").Return(4),
			y:    mockCoordinate.On("Y").Return(4),
		},
		{
			name: "X out of bounds",
			x:    mockCoordinate.On("X").Return(4),
			y:    mockCoordinate.On("Y").Return(3),
		},
		{
			name: "Y out of bounds",
			x:    mockCoordinate.On("X").Return(3),
			y:    mockCoordinate.On("Y").Return(4),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// asserting calls would depend on which axis was beyond size
			assert.True(t, rock.IsBeyond(*sizeLimit))
		})
	}
}
