package coordinate2d_test

import (
	coordinate2d "mars_rover/internal/domain/coordinate/coordinate2D"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoesNotWrapXIfWithinLimit(t *testing.T) {
	coord := coordinate2d.New(1, 1)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.X(), 1)
}

func TestWrapsXIfTooLarge(t *testing.T) {
	coord := coordinate2d.New(3, 1)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.X(), 0)
}

func TestWrapsXIfTooSmall(t *testing.T) {
	coord := coordinate2d.New(-1, 1)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.X(), 2)
}

func TestDoesNotWrapYIfWithinLimit(t *testing.T) {
	coord := coordinate2d.New(1, 1)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.Y(), 1)
}

func TestWrapsYIfTooLarge(t *testing.T) {
	coord := coordinate2d.New(1, 3)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.Y(), 0)
}

func TestWrapsYIfTooSmall(t *testing.T) {
	coord := coordinate2d.New(1, -1)
	limit, _ := size.From(2, 2)

	coord.WrapIfOutOf(*limit)

	assert.Equal(t, coord.Y(), 2)
}

func TestWrapIfOutOf(t *testing.T) {
	limit := size.Size{Width: 2, Height: 2}
	tests := []struct {
		name     string
		starting *coordinate2d.Coordinate2D
		want     *coordinate2d.Coordinate2D
	}{
		{
			name:     "Does not wrap X if within limit",
			starting: coordinate2d.New(1, 1),
			want:     coordinate2d.New(1, 1),
		},
		{
			name:     "Wraps X if too large",
			starting: coordinate2d.New(3, 1),
			want:     coordinate2d.New(0, 1),
		},
		{
			name:     "Wraps X if too small",
			starting: coordinate2d.New(-1, 1),
			want:     coordinate2d.New(2, 1),
		},
		{
			name:     "Does not wrap Y if within limit",
			starting: coordinate2d.New(1, 1),
			want:     coordinate2d.New(1, 1),
		},
		{
			name:     "Wraps Y if too large",
			starting: coordinate2d.New(1, 3),
			want:     coordinate2d.New(1, 0),
		},
		{
			name:     "Wraps Y if too small",
			starting: coordinate2d.New(1, -1),
			want:     coordinate2d.New(1, 2),
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.starting.WrapIfOutOf(limit)
			assert.True(t, testCase.starting.Equals(testCase.want))
		})
	}
}
