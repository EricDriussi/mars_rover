package random_creator

import (
	. "mars_rover/src/action/error"
)

type CreationError struct {
	GameNotCreatedError   *GameNotCreatedError
	GameNotPersistedError *GameNotPersistedError
}

func (e CreationError) Error() string {
	if e.GameNotCreatedError != nil {
		return e.GameNotCreatedError.Error()
	}
	if e.GameNotPersistedError != nil {
		return e.GameNotPersistedError.Error()
	}
	return "unknown error"
}

func (e CreationError) Type() ErrorType {
	if e.GameNotCreatedError != nil {
		return GameNotCreated
	}
	if e.GameNotPersistedError != nil {
		return GameNotPersisted
	}
	return -1
}

func GameNotCreatedErr(err error) *CreationError {
	return &CreationError{
		GameNotCreatedError: &GameNotCreatedError{
			Err: err,
		},
	}
}

func GameNotPersistedErr(err error) *CreationError {
	return &CreationError{
		GameNotPersistedError: &GameNotPersistedError{
			Err: err,
		},
	}
}
