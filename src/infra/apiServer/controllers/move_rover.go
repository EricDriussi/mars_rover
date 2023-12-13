package controllers

import (
	"github.com/google/uuid"
	. "mars_rover/src/action"
	"mars_rover/src/action/command"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

type MoveRequest struct {
	Commands string `json:"commands"`
	Id       string `json:"id"`
}

func MoveRover(action Action, request MoveRequest, responseHandler HTTPResponseHandler) {
	roverId, err := uuid.Parse(request.Id)
	// TODO.LM: this validation would not happen if the uuid was wrapped
	// It's already being validated in the calling controller
	if err != nil {
		responseHandler.SendInternalServerError(err.Error())
		return
	}
	applicationCommands := command.FromString(request.Commands)
	movementResults, actionErr := action.MoveSequence(roverId, applicationCommands)
	if actionErr != nil {
		if actionErr.Type() == RoverNotFound {
			responseHandler.SendBadRequest(actionErr.Error())
			return
		}
		if actionErr.Type() == RoverNotUpdated {
			responseHandler.SendInternalServerError(actionErr.Error())
			return
		}
	}

	responseHandler.SendOk(dto.FromActionResult(movementResults))
}
