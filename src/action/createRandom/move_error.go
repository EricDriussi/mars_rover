package random_creator

import (
	. "github.com/google/uuid"
	. "mars_rover/src/action/error"
)

type CreationError struct {
	RoverNotCreatedError  *RoverNotPersistedError
	PlanetNotCreatedError *PlanetNotPersistedError
	GameNotCreatedError   *GameNotCreatedError
}

func (e CreationError) Error() string {
	if e.RoverNotCreatedError != nil {
		return e.RoverNotCreatedError.Error()
	}
	if e.PlanetNotCreatedError != nil {
		return e.PlanetNotCreatedError.Error()
	}
	if e.GameNotCreatedError != nil {
		return e.GameNotCreatedError.Error()
	}
	return "unknown error"
}

func (e CreationError) Type() ErrorType {
	if e.RoverNotCreatedError != nil {
		return RoverNotPersisted
	}
	if e.PlanetNotCreatedError != nil {
		return PlanetNotPersisted
	}
	if e.GameNotCreatedError != nil {
		return GameNotCreated
	}
	return -1
}

func BuildRoverNotPersistedErr(id UUID, err error) *CreationError {
	return &CreationError{
		RoverNotCreatedError: &RoverNotPersistedError{
			ID:  id,
			Err: err,
		},
	}
}

func BuildPlanetNotPersistedErr(err error) *CreationError {
	return &CreationError{
		PlanetNotCreatedError: &PlanetNotPersistedError{
			Err: err,
		},
	}
}

func BuildGameNotCreatedErr(err error) *CreationError {
	return &CreationError{
		GameNotCreatedError: &GameNotCreatedError{
			Err: err,
		},
	}
}
