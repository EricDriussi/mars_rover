package controllers

import (
	. "mars_rover/src/action"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

func RandomGame(action CreateRandomAction, responseHandler HTTPResponseHandler) {
	curiosity, actionErr := action.Create()
	if actionErr != nil {
		responseHandler.SendInternalServerError(actionErr.Error())
		return
	}

	responseHandler.SendOk(dto.FromDomainRover(curiosity))
	return
}
