package test_helpers

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func contains(stringArray []string, string string) bool {
	for _, element := range stringArray {
		if strings.Contains(element, string) {
			return true
		}
	}
	return false
}

func AssertContains(t *testing.T, stringArray []string, containedString string) bool {
	return assert.True(t, contains(stringArray, containedString))
}
