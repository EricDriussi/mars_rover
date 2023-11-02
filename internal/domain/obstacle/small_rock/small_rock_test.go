package small_rock_test

import (
	"mars_rover/internal/domain/coordinate/test"
	rock "mars_rover/internal/domain/obstacle/small_rock"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIsWithinLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 5, Height: 5}
	mockCoordinate := new(test.MockCoordinate)
	mockCoordinate.On("IsOutsideOf", mock.Anything).Return(false)
	testObstacle := rock.In(mockCoordinate)

	assert.False(t, testObstacle.IsBeyond(*sizeLimit))
	mockCoordinate.AssertCalled(t, "IsOutsideOf", *sizeLimit)
}

func TestIsBeyondLimit(t *testing.T) {
	sizeLimit := &size.Size{Width: 5, Height: 5}
	mockCoordinate := new(test.MockCoordinate)
	mockCoordinate.On("IsOutsideOf", mock.Anything).Return(true)
	testObstacle := rock.In(mockCoordinate)

	assert.True(t, testObstacle.IsBeyond(*sizeLimit))
	mockCoordinate.AssertCalled(t, "IsOutsideOf", *sizeLimit)
}
