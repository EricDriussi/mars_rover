package load

import (
	"fmt"
)

type errorType int

const (
	gameNotLoaded errorType = iota
	gameNotFound
)

type LoadError struct {
	errType errorType
	errMsg  string
}

func (e LoadError) Error() string {
	if e.errType == gameNotLoaded {
		return fmt.Sprintf("could not load game: %s", e.errMsg)
	}
	return "unknown error"
}

func (e LoadError) IsNotFound() bool {
	return e.errType == gameNotFound
}

func GameNotLoaded(err error) *LoadError {
	return &LoadError{
		errType: gameNotLoaded,
		errMsg:  err.Error(),
	}
}

func GameNotFound(err error) *LoadError {
	return &LoadError{
		errType: gameNotFound,
		errMsg:  err.Error(),
	}
}
