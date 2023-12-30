package test_helpers

import "errors"

func SuccessfulRoverFunc() func() error {
	return func() error {
		return nil
	}
}

func RoverFunc(fn func() error) func() error {
	return func() error {
		return fn()
	}
}

func FailedRoverFunc() func() error {
	return func() error {
		return errors.New("an error")
	}
}
