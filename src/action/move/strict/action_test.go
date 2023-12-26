package strict_mover_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/mock"
	. "mars_rover/src/action/move/command"
	"mars_rover/src/action/move/strict"
	"mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/test_helpers"
	"mars_rover/src/test_helpers/mocks"
	. "mars_rover/src/test_helpers/mocks"
	"testing"
)

func TestDoesNotAbortASingleSuccessfulMovementCommand(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Forward}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "MoveForward")
	AssertEncounteredNoIssues(t, movementResults)
	assert.Len(t, movementResults, len(commands))
}

func TestDoesNotAbortASingleRotationCommand(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "TurnRight")
	AssertEncounteredNoIssues(t, movementResults)
	assert.Len(t, movementResults, len(commands))
}

func TestAbortsASingleFailedMovementCommand(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	testRover.On("MoveForward").Return(errors.New("movement blocked"))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Forward}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "MoveForward")
	AssertEncounteredAnIssue(t, movementResults)
	assert.Len(t, movementResults, len(commands))
}

func TestCallsRoverBasedOnSuccessfulCommandsInGivenSequence(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Backward, Forward, Left, Right}
	_, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "MoveBackward")
	testRover.AssertCalled(t, "MoveForward")
	testRover.AssertCalled(t, "TurnLeft")
	testRover.AssertCalled(t, "TurnRight")
}

func TestReportsRoverMovementBasedOnSuccessfulCommandsInGivenSequence(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Backward, Forward, Left, Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	assert.Len(t, movementResults, len(commands))
	AssertContainsOrderedCommands(t, movementResults, commands)
	AssertEncounteredNoIssues(t, movementResults)
}

func TestStopsCallingRoverOnceACommandFails(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	testRover.On("MoveBackward").Return(nil)
	testRover.On("MoveForward").Return(errors.New("movement blocked"))
	testRover.On("TurnLeft").Return()
	testRover.On("TurnRight").Return()
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Backward, Forward, Left, Right}
	_, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	testRover.AssertCalled(t, "MoveBackward")
	testRover.AssertCalled(t, "MoveForward")
	testRover.AssertNotCalled(t, "TurnLeft")
	testRover.AssertNotCalled(t, "TurnRight")
}

func TestReportsRoverMovementUntilACommandFails(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	testRover.On("MoveBackward").Return(nil)
	testRover.On("MoveForward").Return(errors.New("movement blocked"))
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(nil)

	act := strict_mover.With(repo)
	commands := Commands{Backward, Forward, Left, Right}
	movementResults, err := act.Move(uuid.New(), commands)

	assert.Nil(t, err)
	assert.Len(t, movementResults, len(commands)-2)
	AssertContainsOrderedCommands(t, movementResults, Commands{Backward, Forward})
	AssertEncounteredAnIssue(t, movementResults)
}

func TestReportsRepoErrorWhenGettingRover(t *testing.T) {
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(new(MockRover), errors.New("whatever"))

	act := strict_mover.With(repo)
	irrelevantCommand := Forward
	movementResults, err := act.Move(uuid.New(), Commands{irrelevantCommand})

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestReportsRepoErrorWhenUpdatingRover(t *testing.T) {
	testRover := mocks.LandedRover(*absoluteCoordinate.From(1, 1))
	mocks.MakeAlwaysSuccessful(testRover)
	repo := new(MockRepo)
	repo.On("GetRover", Anything).Return(testRover, nil)
	repo.On("UpdateRover").Return(errors.New("whatever"))

	act := strict_mover.With(repo)
	irrelevantCommand := Forward
	movementResults, err := act.Move(uuid.New(), Commands{irrelevantCommand})

	assert.Empty(t, movementResults)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "failed to update")
}
