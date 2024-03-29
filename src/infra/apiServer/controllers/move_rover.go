package controllers

import (
	. "mars_rover/src/action"
	. "mars_rover/src/action/mover"
	"mars_rover/src/action/mover/command"
	"mars_rover/src/domain/rover/id"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

type MoveRequest struct {
	Commands string `json:"commands"`
	Id       string `json:"id"`
}

func MoveRover(action MoveAction, request MoveRequest, responseHandler HTTPResponseHandler) {
	roverId, err := id.Parse(request.Id)
	if err != nil {
		responseHandler.SendBadRequest("Invalid ID")
		return
	}
	applicationCommands := command.FromString(request.Commands)
	if len(applicationCommands) == 0 {
		responseHandler.SendBadRequest("No valid commands provided")
		return
	}

	movementResults, actionErr := action.Move(roverId, applicationCommands)
	if actionErr != nil {
		sendErrorResponseBasedOnType(actionErr, responseHandler)
		return
	}

	responseHandler.SendOk(dto.FromMovementResults(movementResults))
}

func sendErrorResponseBasedOnType(actionErr *MovementError, responseHandler HTTPResponseHandler) {
	if actionErr.IsNotFound() {
		responseHandler.SendBadRequest(actionErr.Error())
	} else {
		responseHandler.SendInternalServerError(actionErr.Error())
	}
}
