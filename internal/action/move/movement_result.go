package move

import . "mars_rover/internal/domain/rover"

type MovementResult struct {
	Rover          Rover
	MovementErrors *MovementErrors
	Error          error
}

func (this *MovementResult) HasMovementErrors() bool {
	return this.MovementErrors != nil && len(this.MovementErrors.errors) > 0
}
