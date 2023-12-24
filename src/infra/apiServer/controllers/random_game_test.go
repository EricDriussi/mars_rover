package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/createRandom"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/obstacle"
	"mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover/direction"
	"mars_rover/src/domain/rover/planetMap"
	"mars_rover/src/domain/size"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestSendsOkResponseWhenCreateRandomActionIsSuccessful(t *testing.T) {
	testSize, _ := size.Square(10)
	mockPlanet := new(MockPlanet)
	mockPlanet.On("Size").Return(*testSize)
	mockObstacle := new(MockObstacle)
	mockObstacle.On("Occupies", Anything).Return(false)
	mockObstacle.On("Coordinates").Return([]absoluteCoordinate.AbsoluteCoordinate{})
	testObstacles := obstacles.FromList([]obstacle.Obstacle{mockObstacle})
	mockPlanet.On("Obstacles").Return(*testObstacles)
	mockRover := new(MockRover)
	mockRover.On("Id").Return(uuid.New())
	mockRover.On("Direction").Return(North{})
	mockRover.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	mockRover.On("Map").Return(*planetMap.OfPlanet(mockPlanet))
	mockAction := new(MockAction)
	mockAction.On("Create").Return(mockRover, nil)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendOk", Anything).Return()

	RandomGame(mockAction, mockHandler)

	mockHandler.AssertCalled(t, "SendOk", Anything)
}

func TestSendsInternalServerErrorResponseWhenCreateRandomActionReportsAnError(t *testing.T) {
	mockAction := new(MockAction)
	mockRover := new(MockRover)
	creationError := BuildGameNotCreatedErr(errors.New("test error"))
	mockAction.On("Create").Return(mockRover, creationError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendInternalServerError", Anything).Return()

	RandomGame(mockAction, mockHandler)

	mockHandler.AssertCalled(t, "SendInternalServerError", Anything)
}
