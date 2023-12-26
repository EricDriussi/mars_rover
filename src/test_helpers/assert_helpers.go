package test_helpers

import (
	"github.com/stretchr/testify/assert"
	"mars_rover/src/action"
	. "mars_rover/src/action/move/command"
	"testing"
)

func AssertEncounteredNoIssues(t *testing.T, results []action.MovementResult) {
	for _, result := range results {
		assert.False(t, result.IssueDetected)
	}
}

func AssertEncounteredAnIssue(t *testing.T, results []action.MovementResult) {
	assert.True(t, containsAnIssue(results))
}

func containsAnIssue(results []action.MovementResult) bool {
	for _, result := range results {
		if result.IssueDetected {
			return true
		}
	}
	return false
}

func AssertContainsOrderedCommands(t *testing.T, movementResults []action.MovementResult, commands Commands) {
	assert.Len(t, movementResults, len(commands))
	for i, cmd := range commands {
		assert.Equal(t, movementResults[i].Cmd.String(), cmd.String())
	}
}
