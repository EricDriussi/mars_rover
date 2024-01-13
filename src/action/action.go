package action

import (
	. "mars_rover/src/action/gameLoader"
	. "mars_rover/src/action/mover"
	. "mars_rover/src/action/mover/command"
	. "mars_rover/src/action/randomCreator"
	. "mars_rover/src/domain"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/id"
)

type CreateRandomAction interface {
	// TODO: Also return planet (color)?
	Create() (Rover, *CreationError)
}

type MoveAction interface {
	Move(roverId ID, commands Commands) ([]MovementResult, *MovementError)
	// TODO.LM: here I'm returning a result AND an error
	// I understand this is strange to see, but it is in line with
	// how error handling usually works in Go
	// I'm not sure if this is better thant wrapping the error inside the MovementResult
}

type LoadAction interface {
	Load(roverId ID) (*Game, *LoadError)
}
