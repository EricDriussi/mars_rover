package emptyPlanet_test

import (
	"mars_rover/internal/domain/obstacle"
	"mars_rover/internal/domain/planet/emptyPlanet"
	"mars_rover/internal/domain/size"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetsSize(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	planet, _ := emptyPlanet.Create("testColor", *sizeLimit)

	assert.Equal(t, *sizeLimit, planet.Size())
}

func TestGetsObstacles(t *testing.T) {
	sizeLimit, _ := size.Square(5)
	planet, _ := emptyPlanet.Create("testColor", *sizeLimit)

	assert.Equal(t, []obstacle.Obstacle{}, planet.Obstacles())
}
