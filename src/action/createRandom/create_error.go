package random_creator

import (
	. "mars_rover/src/action/error"
)

type CreationError struct {
	GameNotPersistedError *GameNotPersistedError
}

func (e CreationError) Error() string {
	if e.GameNotPersistedError != nil {
		return e.GameNotPersistedError.Error()
	}
	return "unknown error"
}

func (e CreationError) Type() ErrorType {
	if e.GameNotPersistedError != nil {
		return GameNotPersisted
	}
	return -1
}

func GameNotPersistedErr(err error) *CreationError {
	return &CreationError{
		GameNotPersistedError: &GameNotPersistedError{
			Err: err,
		},
	}
}
