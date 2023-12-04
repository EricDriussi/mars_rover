package move

import (
	"fmt"
)

type MovementErrors struct {
	errors []MovementError
}

func (this *MovementErrors) Add(command string, err error) {
	this.errors = append(this.errors, MovementError{command: command, err: err})
}

func (this *MovementErrors) List() []MovementError {
	return this.errors
}

func (this *MovementErrors) AsStringArray() []string {
	var errors []string
	for _, err := range this.errors {
		errors = append(errors, err.AsString())
	}
	return errors
}

type MovementError struct {
	command string
	err     error
}

func (this *MovementError) AsString() string {
	return fmt.Sprintf("error executing command %v: %v", this.command, this.err)
}
