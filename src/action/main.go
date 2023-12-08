package action

import (
	. "github.com/google/uuid"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
)

type Action interface {
	Random() (Rover, error)
	MoveSequence(roverId UUID, commands string) MovementResult
}

type LaxAction struct {
	repo Repository
}

func For(repo Repository) *LaxAction {
	return &LaxAction{
		repo: repo,
	}
}
