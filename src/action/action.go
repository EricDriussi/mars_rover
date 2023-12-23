package action

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
)

type CreateRandomAction interface {
	Create() (Rover, *CreationError)
}

type MoveAction interface {
	Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError)
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
