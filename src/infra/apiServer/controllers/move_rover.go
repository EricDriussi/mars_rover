package controllers

import (
	"github.com/google/uuid"
	. "mars_rover/src/action"
	. "mars_rover/src/action/error"
	"mars_rover/src/action/move"
	"mars_rover/src/action/move/command"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

type MoveRequest struct {
	Commands string `json:"commands"`
	Id       string `json:"id"`
}

func MoveRover(action MoveAction, request MoveRequest, responseHandler HTTPResponseHandler) {
	roverId, err := uuid.Parse(request.Id)
	// TODO.LM: this validation would not happen if the uuid was wrapped
	// It's already being validated in the calling controller
	if err != nil {
		responseHandler.SendInternalServerError(err.Error())
		return
	}
	applicationCommands := command.FromString(request.Commands)
	if len(applicationCommands) == 0 {
		responseHandler.SendBadRequest("No valid commands provided")
		return
	}

	movementResults, actionErr := action.Move(roverId, applicationCommands)
	if actionErr != nil {
		sendResponseBasedOnErrorType(actionErr, responseHandler)
		return
	}

	responseHandler.SendOk(dto.FromMovementResult(movementResults))
}

func sendResponseBasedOnErrorType(actionErr *move.MovementError, responseHandler HTTPResponseHandler) {
	if actionErr.Type() == RoverNotFound {
		responseHandler.SendBadRequest(actionErr.Error())
	}
	if actionErr.Type() == RoverNotUpdated {
		responseHandler.SendInternalServerError(actionErr.Error())
	}
}
