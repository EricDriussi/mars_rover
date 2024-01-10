package action

import (
	. "mars_rover/src/action/createRandom"
	. "mars_rover/src/action/load"
	. "mars_rover/src/action/move"
	. "mars_rover/src/action/move/command"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/uuid"
)

type CreateRandomAction interface {
	// TODO: Also return planet (color)?
	Create() (Rover, *CreationError)
}

type MoveAction interface {
	Move(roverId UUID, commands Commands) ([]MovementResult, *MovementError)
	// TODO.LM: here I'm returning a result AND an error
	// I understand this is strange to see, but it is in line with
	// how error handling usually works in Go
	// I'm not sure if this is better thant wrapping the error inside the MovementResult
}

type LoadAction interface {
	Load(roverId UUID) (*Game, *LoadError)
}
