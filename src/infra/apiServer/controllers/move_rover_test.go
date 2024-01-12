package controllers_test

import (
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/move"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestSendsOkResponseWhenMovementActionIsSuccessful(t *testing.T) {
	mockAction := new(MockAction)
	mockAction.On("Move").Return([]MovementResult{}, nil)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendOk", Anything).Return()

	controllers.MoveRover(mockAction, aMoveRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendOk", Anything)
}

func TestSendsBadRequestResponseWhenNoValidCommandsAreProvided(t *testing.T) {
	mockAction := new(MockAction)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendBadRequest", Anything).Return()
	testCases := []struct {
		name            string
		invalidCommands string
	}{
		{
			name:            "no commands",
			invalidCommands: "",
		},
		{
			name:            "invalid commands",
			invalidCommands: "xxx",
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testRequest := MoveRequest{
				Commands: testCase.invalidCommands,
				Id:       id.New().String(),
			}

			controllers.MoveRover(mockAction, testRequest, mockHandler)

			mockHandler.AssertCalled(t, "SendBadRequest", Anything)
		})
	}
}

func TestSendsBadRequestResponseWhenActionDoesNotFindRover(t *testing.T) {
	notFoundError := NotFoundErr()
	mockAction := new(MockAction)
	mockAction.On("Move").Return([]MovementResult{}, notFoundError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendBadRequest", Anything).Return()

	controllers.MoveRover(mockAction, aMoveRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendBadRequest", Anything)
}

func TestSendsInternalServerErrorResponseWhenActionCannotUpdateRover(t *testing.T) {
	notUpdatedError := NotUpdatedErr()
	mockAction := new(MockAction)
	mockAction.On("Move").Return([]MovementResult{}, notUpdatedError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendInternalServerError", Anything).Return()

	controllers.MoveRover(mockAction, aMoveRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendInternalServerError", Anything)
}

func aMoveRequest() MoveRequest {
	return MoveRequest{
		Commands: "whatever",
		Id:       id.New().String(),
	}
}
