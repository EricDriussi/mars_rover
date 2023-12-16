package action

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
)

// TODO: one action per use case?
type Action interface {
	Random() (Rover, error)
	MoveSequence(roverId UUID, commands Commands) ([]MovementResult, *ActionError)
	// TODO.LM: here I'm returning a result AND an error
	// I understand this is strange to see, but it is in line with
	// how error handling usually works in Go
	// I'm not sure if this is better thant wrapping the error inside the MovementResult
}

type MovementResult struct {
	Cmd           Command
	IssueDetected bool
	Coord         AbsoluteCoordinate
	Dir           Direction
}

type LaxAction struct {
	repo Repository
}

func For(repo Repository) *LaxAction {
	return &LaxAction{
		repo: repo,
	}
}
