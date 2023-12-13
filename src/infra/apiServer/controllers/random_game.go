package controllers

import (
	. "mars_rover/src/action"
	"mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/apiServer/dto"
)

func RandomGame(action Action) (CreateResponseDTO, error) {
	curiosity, err := action.Random()
	if err != nil {
		return CreateResponseDTO{}, err
	}

	return dto.FromDomainRover(curiosity), nil
}
