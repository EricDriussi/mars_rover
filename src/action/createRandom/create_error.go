package randomCreator

import (
	"fmt"
)

type errorType int

const (
	gameNotPersisted errorType = iota
)

type CreationError struct {
	errType errorType
	errMsg  string
}

func (e CreationError) Error() string {
	if e.errType == gameNotPersisted {
		return fmt.Sprintf("could not persist game: %s", e.errMsg)
	}
	return "unknown error"
}

func BuildGameNotPersistedErr(err error) *CreationError {
	return &CreationError{
		errType: gameNotPersisted,
		errMsg:  err.Error(),
	}
}
