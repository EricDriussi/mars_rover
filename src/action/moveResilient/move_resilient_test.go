package moveResilient_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/command"
	"mars_rover/src/action/moveResilient"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
	. "mars_rover/src/test_helpers"
	. "mars_rover/src/test_helpers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlesASingleSuccessfulMovementCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveForward").Return(nil)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Forward}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
	AssertEncounteredNoIssues(t, movementResults)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveForward").Return(errors.New("an error"))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Forward}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
	AssertEncounteredAnIssue(t, movementResults)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "TurnRight")
	AssertEncounteredNoIssues(t, movementResults)
}

func TestCallsRoverBasedOnSuccessfulCommandsWithGivenSequence(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("TurnLeft").Return()
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Backward, Forward, Left, Right}
	_, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "TurnRight")
}

func TestReportsRoverMovementBasedOnSuccessfulCommandsWithGivenSequence(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(nil)
	curiosity.On("TurnLeft").Return()
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Backward, Forward, Left, Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	assert.Len(t, movementResults, len(commands))
	AssertContainsOrderedCommands(t, movementResults, commands)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestCallsRoverBasedOnSuccessfulAndFailedCommandsWithGivenSequence(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New("movement blocked"))
	curiosity.On("TurnLeft").Return()
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Backward, Forward, Left, Right}
	_, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "TurnLeft")
	curiosity.AssertCalled(t, "TurnRight")
}

func TestReportsRoverMovementBasedOnSuccessfulAndFailedCommandsWithGivenSequence(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New("movement blocked"))
	curiosity.On("TurnLeft").Return()
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := moveResilient.For(repo)
	commands := Commands{Backward, Forward, Left, Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	assert.Len(t, movementResults, len(commands))
	AssertContainsOrderedCommands(t, movementResults, commands)
	AssertEncounteredAnIssue(t, movementResults)
}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(new(MockRover), errors.New("whatever"))

	act := moveResilient.For(repo)
	irrelevantCommand := Forward
	movementResults, err := act.Move(uuid.New(), Commands{irrelevantCommand})

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestReportsRepoErrorWhenUpdatingRover(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveForward").Return(nil)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(errors.New("whatever"))

	act := moveResilient.For(repo)
	irrelevantCommand := Forward
	movementResults, err := act.Move(uuid.New(), Commands{irrelevantCommand})

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update")
}
