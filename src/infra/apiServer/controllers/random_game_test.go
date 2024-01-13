package controllers_test

import (
	"errors"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/randomCreator"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/infra/apiServer/controllers"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestSendsOkResponseWhenCreateRandomActionIsSuccessful(t *testing.T) {
	mockPlanet := mocks.PlanetWithNoObstaclesOfSize(t, 10)
	mockRover := mocks.RoverIn(mockPlanet, *absoluteCoordinate.Build(1, 1))
	mockAction := new(MockAction)
	mockAction.On("Create").Return(mockRover, nil)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendOk", Anything).Return()

	controllers.RandomGame(mockAction, mockHandler)

	mockHandler.AssertCalled(t, "SendOk", Anything)
}

func TestSendsInternalServerErrorResponseWhenCreateRandomActionReportsAnError(t *testing.T) {
	mockRover := new(MockRover)
	mockAction := new(MockAction)
	creationError := GameNotPersistedErr(errors.New("test error"))
	mockAction.On("Create").Return(mockRover, creationError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendInternalServerError", Anything).Return()

	controllers.RandomGame(mockAction, mockHandler)

	mockHandler.AssertCalled(t, "SendInternalServerError", Anything)
}
