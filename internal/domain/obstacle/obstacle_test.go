package obstacle_test

import (
	"mars_rover/internal/domain/coordinate"
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 5, Height: 5}
	mockCoordinate := new(MockCoordinate)
	mockCoordinate.On("IsOutsideOf", mock.Anything).Return(false)
	testObstacle := obstacle.In(mockCoordinate)

	assert.False(t, testObstacle.IsBeyond(*sizeLimit))
	mockCoordinate.AssertCalled(t, "IsOutsideOf", *sizeLimit)
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 5, Height: 5}
	mockCoordinate := new(MockCoordinate)
	mockCoordinate.On("IsOutsideOf", mock.Anything).Return(true)
	testObstacle := obstacle.In(mockCoordinate)

	assert.True(t, testObstacle.IsBeyond(*sizeLimit))
	mockCoordinate.AssertCalled(t, "IsOutsideOf", *sizeLimit)
}

type MockCoordinate struct {
	mock.Mock
	x, y int
}

func (c *MockCoordinate) WrapIfOutOf(limit size.Size) {
}

func (c *MockCoordinate) IsOutsideOf(limit size.Size) bool {
	args := c.Called(limit)
	return args.Bool(0)
}

func (c *MockCoordinate) Equals(other coordinate.Coordinate) bool {
	args := c.Called(other)
	return args.Bool(0)
}

func (c *MockCoordinate) X() int {
	return c.x
}

func (c *MockCoordinate) Y() int {
	return c.y
}
