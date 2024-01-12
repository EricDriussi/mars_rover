package domain

import (
	"fmt"
)

type errorType int

const (
	persistenceMalfunction errorType = iota
	couldNotUpdate
	couldNotMap
	couldNotAdd
	notFound
)

type RepositoryError struct {
	errType errorType
	errMsg  string
}

func (this RepositoryError) Error() string {
	if this.errType == persistenceMalfunction {
		return fmt.Sprintf("persistence malfunction: %s", this.errMsg)
	}
	if this.errType == couldNotUpdate {
		return fmt.Sprintf("could not update resource: %s", this.errMsg)
	}
	if this.errType == notFound {
		return this.errMsg
	}
	if this.errType == couldNotMap {
		return fmt.Sprintf("could not map resource: %s", this.errMsg)
	}
	if this.errType == couldNotAdd {
		return fmt.Sprintf("could not add resource: %s", this.errMsg)
	}
	return "unknown error"
}

func (this RepositoryError) IsNotFound() bool {
	return this.errType == notFound
}

func PersistenceMalfunction(err error) *RepositoryError {
	return &RepositoryError{
		errType: persistenceMalfunction,
		errMsg:  err.Error(),
	}
}

func CouldNotUpdate(err error) *RepositoryError {
	return &RepositoryError{
		errType: persistenceMalfunction,
		errMsg:  err.Error(),
	}
}

func NotFound() *RepositoryError {
	return &RepositoryError{
		errType: notFound,
		errMsg:  "resource not found",
	}
}

func CouldNotMap(err error) *RepositoryError {
	return &RepositoryError{
		errType: couldNotMap,
		errMsg:  err.Error(),
	}
}

func CouldNotAdd(err error) *RepositoryError {
	return &RepositoryError{
		errType: couldNotAdd,
		errMsg:  err.Error(),
	}
}
