package load

import (
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover/uuid"
)

type SimpleGameLoader struct {
	repo Repository
}

func With(repo Repository) *SimpleGameLoader {
	return &SimpleGameLoader{
		repo: repo,
	}
}

func (this *SimpleGameLoader) Load(roverId UUID) (*Game, *LoadError) {
	game, getErr := this.repo.GetGame(roverId)
	if getErr != nil {
		if getErr.IsNotFound() {
			return nil, GameNotFound(getErr)
		}
		return nil, GameNotLoaded(getErr)
	}

	return game, nil
}
