package action_test

import (
	"errors"
	"github.com/google/uuid"
	. "github.com/stretchr/testify/mock"
	"mars_rover/src/action"
	"mars_rover/src/action/command"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/rover/direction"
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

	act := action.For(repo)
	commands := command.Commands{command.Forward}
	movementResults, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
	assertEncounteredNoIssues(t, movementResults)
}

func TestHandlesASingleFailedMovementCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveForward").Return(errors.New("an error"))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := action.For(repo)
	commands := command.Commands{command.Forward}
	movementResults, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveForward")
	assertEncounteredAnIssue(t, movementResults)
}

func TestHandlesASingleTurningCommand(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := action.For(repo)
	commands := command.Commands{command.Right}
	movementResults, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "TurnRight")
	assertEncounteredNoIssues(t, movementResults)
}

func TestHandlesACombinationOfSuccessfulAndUnsuccessfulCommands(t *testing.T) {
	curiosity := new(MockRover)
	curiosity.On("Coordinate").Return(*absoluteCoordinate.From(1, 1))
	curiosity.On("Direction").Return(North{})
	curiosity.On("MoveBackward").Return(nil)
	curiosity.On("MoveForward").Return(errors.New("movement blocked"))
	curiosity.On("TurnLeft").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(curiosity, nil)
	repo.On("UpdateRover").Return(nil)

	act := action.For(repo)
	commands := command.Commands{command.Backward, command.Forward, command.Left}
	movementResults, err := act.MoveSequence(uuid.New(), commands)

	assert.Nil(t, err)
	curiosity.AssertCalled(t, "MoveBackward")
	curiosity.AssertCalled(t, "MoveForward")
	curiosity.AssertCalled(t, "TurnLeft")
	firstCommandResult := movementResults[0]
	assert.Equal(t, firstCommandResult.Cmd.ToString(), "b")
	assert.False(t, firstCommandResult.IssueDetected)
	secondCommandResult := movementResults[1]
	assert.Equal(t, secondCommandResult.Cmd.ToString(), "f")
	assert.True(t, secondCommandResult.IssueDetected)
	thirdCommandResult := movementResults[2]
	assert.Equal(t, thirdCommandResult.Cmd.ToString(), "l")
	assert.False(t, thirdCommandResult.IssueDetected)

}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(new(MockRover), errors.New("whatever"))

	act := action.For(repo)
	irrelevantCommand := command.Forward
	movementResults, err := act.MoveSequence(uuid.New(), command.Commands{irrelevantCommand})

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

	act := action.For(repo)
	irrelevantCommand := command.Forward
	movementResults, err := act.MoveSequence(uuid.New(), command.Commands{irrelevantCommand})

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update")
}

func assertEncounteredNoIssues(t *testing.T, result []action.MovementResult) {
	for _, res := range result {
		assert.False(t, res.IssueDetected)
	}
}

func assertEncounteredAnIssue(t *testing.T, result []action.MovementResult) {
	for _, res := range result {
		assert.True(t, res.IssueDetected)
	}
}
