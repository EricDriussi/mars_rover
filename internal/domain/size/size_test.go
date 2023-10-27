package size_test

import (
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotAllowNegativeValues(t *testing.T) {
	testCases := []struct {
		name   string
		width  int
		height int
	}{
		{
			name:   "neither width nor height can be negative",
			width:  10,
			height: -10,
		},
		{
			name:   "width cannot be negative",
			width:  -10,
			height: 10,
		},
		{
			name:   "height cannot be negative",
			width:  10,
			height: -10,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := size.From(testCase.width, testCase.height)
			assert.Error(t, err)
		})
	}
}

func TestDoesNotAllowZeroValues(t *testing.T) {
	testCases := []struct {
		name   string
		width  int
		height int
	}{
		{
			name:   "neither width nor height can be zero",
			width:  0,
			height: 0,
		},
		{
			name:   "width cannot be zero",
			width:  0,
			height: 0,
		},
		{
			name:   "height cannot be zero",
			width:  0,
			height: 0,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := size.From(testCase.width, testCase.height)
			assert.Error(t, err)
		})
	}
}
