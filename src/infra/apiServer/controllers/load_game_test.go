package controllers_test

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/domain"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/apiServer/controllers"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestSendsOkResponseWhenLoadActionIsSuccessful(t *testing.T) {
	mockPlanet := mocks.PlanetWithNoObstaclesOfSize(t, 10)
	mockRover := mocks.RoverIn(mockPlanet, *absoluteCoordinate.Build(1, 1))
	gameDTO := Game{
		Rover:  mockRover,
		Planet: mockPlanet,
	}
	mockAction := new(MockAction)
	mockAction.On("Load").Return(&gameDTO, nil)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendOk", Anything).Return()

	controllers.LoadGame(mockAction, aLoadRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendOk", Anything)
}

func aLoadRequest() LoadRequest {
	return LoadRequest{
		Id: id.New().String(),
	}
}
