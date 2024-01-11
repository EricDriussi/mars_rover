package controllers

import (
	. "mars_rover/src/action"
	. "mars_rover/src/action/load"
	"mars_rover/src/domain/rover/uuid"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

type LoadRequest struct {
	Id string `json:"id"`
}

func LoadGame(action LoadAction, request LoadRequest, responseHandler HTTPResponseHandler) {
	roverId, err := uuid.Parse(request.Id)
	if err != nil {
		responseHandler.SendBadRequest("Invalid ID")
		return
	}

	game, actionErr := action.Load(roverId)
	if actionErr != nil {
		sendLoadErrorBasedOnType(actionErr, responseHandler)
		return
	}

	responseHandler.SendOk(dto.FromGame(game))
}

func sendLoadErrorBasedOnType(actionErr *LoadError, responseHandler HTTPResponseHandler) {
	if actionErr.IsNotFound() {
		responseHandler.SendBadRequest(actionErr.Error())
	} else {
		responseHandler.SendInternalServerError(actionErr.Error())
	}
}
