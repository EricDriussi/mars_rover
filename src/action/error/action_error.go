package error

import (
	"fmt"
	. "github.com/google/uuid"
)

type ErrorType int

// TODO: do these belong to the repo??
// is GameNotCreated the only real action error?
const (
	RoverNotFound ErrorType = iota
	RoverNotUpdated
	RoverNotPersisted
	PlanetNotPersisted
	GameNotCreated
)

type RoverNotFoundError struct {
	ID  UUID
	Err error
}

func (e RoverNotFoundError) Error() string {
	return fmt.Sprintf("rover with ID %s not found", e.ID)
}

type RoverNotUpdatedError struct {
	ID  UUID
	Err error
}

func (e RoverNotUpdatedError) Error() string {
	return fmt.Sprintf("failed to update rover with ID %s", e.ID)
}

type RoverNotPersistedError struct {
	ID  UUID
	Err error
}

func (e RoverNotPersistedError) Error() string {
	return fmt.Sprintf("rover with ID %s could not be saved", e.ID)
}

type PlanetNotPersistedError struct {
	Err error
}

func (e PlanetNotPersistedError) Error() string {
	return "could not save planet"
}

type GameNotCreatedError struct {
	Err error
}

func (e GameNotCreatedError) Error() string {
	return "could not create game"
}
