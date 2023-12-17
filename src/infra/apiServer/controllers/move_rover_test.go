package controllers_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	"mars_rover/src/action"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestSendsOkResponseWhenMovementActionIsSuccessful(t *testing.T) {
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]action.MovementResult{}, nil)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendOk", Anything).Return()

	MoveRover(mockAction, aMoveRequest(), mockHandler)

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
		
			MoveRover(mockAction,
				MoveRequest{
					Commands: testCase.invalidCommands,
					Id:       uuid.New().String(),
				},
				mockHandler)

			mockHandler.AssertCalled(t, "SendBadRequest", Anything)
		})
	}
}

func TestSendsBadRequestResponseWhenActionDoesNotFindRover(t *testing.T) {
	notFoundError := action.BuildNotFound(uuid.New(), errors.New("whatever"))
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]action.MovementResult{}, notFoundError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendBadRequest", Anything).Return()

	MoveRover(mockAction, aMoveRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendBadRequest", Anything)
}

func TestSendsInternalServerErrorResponseWhenActionCannotUpdateRover(t *testing.T) {
	notUpdatedError := action.BuildNotUpdated(uuid.New(), errors.New("whatever"))
	mockAction := new(MockAction)
	mockAction.On("MoveSequence").Return([]action.MovementResult{}, notUpdatedError)
	mockHandler := new(MockHTTPResponseHandler)
	mockHandler.On("SendInternalServerError", Anything).Return()

	MoveRover(mockAction, aMoveRequest(), mockHandler)

	mockHandler.AssertCalled(t, "SendInternalServerError", Anything)
}

func aMoveRequest() MoveRequest {
	return MoveRequest{
		Commands: "whatever",
		Id:       uuid.New().String(),
	}
}
