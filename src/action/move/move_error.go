package move

import (
	"fmt"
	. "mars_rover/src/domain/rover/id"
)

type errorType int

const (
	roverNotFound errorType = iota
	roverNotUpdated
)

type MovementError struct {
	ID      ID
	errType errorType
}

func (e MovementError) Error() string {
	if e.IsNotFound() {
		return fmt.Sprintf("rover with ID %s not found", e.ID)
	}
	if e.IsNotUpdated() {
		return fmt.Sprintf("failed to update rover with ID %s", e.ID)
	}
	return "unknown movement error"
}

func (e MovementError) IsNotFound() bool {
	return e.errType == roverNotFound
}

func (e MovementError) IsNotUpdated() bool {
	return e.errType == roverNotUpdated
}

func NotFoundErr() *MovementError {
	return &MovementError{
		errType: roverNotFound,
	}
}

func NotUpdatedErr() *MovementError {
	return &MovementError{
		errType: roverNotUpdated,
	}
}
