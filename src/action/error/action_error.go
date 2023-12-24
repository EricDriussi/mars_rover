package error

import (
	"fmt"
	. "github.com/google/uuid"
)

type ErrorType int

// TODO: do these belong to the repo??
const (
	RoverNotFound ErrorType = iota
	RoverNotUpdated
	GameNotCreated
	GameNotPersisted
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

type GameNotCreatedError struct {
	Err error
}

func (e GameNotCreatedError) Error() string {
	return fmt.Sprintf("could not create game: %s", e.Err.Error())
}

type GameNotPersistedError struct {
	Err error
}

func (e GameNotPersistedError) Error() string {
	return fmt.Sprintf("could not persist game: %s", e.Err.Error())
}
