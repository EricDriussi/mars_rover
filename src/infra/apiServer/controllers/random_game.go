package controllers

import (
	. "mars_rover/src/action"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

func RandomGame(action Action, responseHandler HTTPResponseHandler) {
	curiosity, err := action.Random()
	if err != nil {
		responseHandler.SendInternalServerError(err.Error())
		return
	}

	responseHandler.SendOk(dto.FromDomainRover(curiosity))
	return
}
