package mover

import (
	. "mars_rover/src/action/mover/command"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/domain/rover/direction"
)

type MovementResult struct {
	Cmd           Command
	IssueDetected bool
	Coord         AbsoluteCoordinate
	Dir           Direction
}

func BuildSuccessfulMovementResult(rover Rover, cmd Command) MovementResult {
	return buildMovementResult(rover, cmd, false)
}

func BuildFailedMovementResult(rover Rover, cmd Command) MovementResult {
	return buildMovementResult(rover, cmd, true)
}

func buildMovementResult(rover Rover, cmd Command, hadIssue bool) MovementResult {
	return MovementResult{
		Cmd:           cmd,
		IssueDetected: hadIssue,
		Coord:         rover.Coordinate(),
		Dir:           rover.Direction(),
	}
}
