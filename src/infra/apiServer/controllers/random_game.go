package controllers

import (
	. "mars_rover/src/action"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/responses"
)

func RandomGame(action CreateRandomAction, responseHandler HTTPResponseHandler) {
	curiosity, err := action.Create()
	if err != nil {
		responseHandler.SendInternalServerError(err.Error())
		return
	}

	responseHandler.SendOk(dto.FromDomainRover(curiosity))
	return
}
