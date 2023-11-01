package planetMap_test

// import (
// 	"fmt"
// 	"mars_rover/internal/domain/coordinate"
// 	"mars_rover/internal/domain/obstacle"
// 	"mars_rover/internal/domain/planet"
// 	"mars_rover/internal/domain/rover/planet_map"
// 	"mars_rover/internal/domain/size"
// 	"math/rand"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestReportsCollisionWithMock(t *testing.T) {
// 	x := 3
// 	y := 3
// 	planet := createPlanetWithObstacleIn(x, y)
// 	planetMap := &planetMap.Map{size: &size.Size{Width: x, Height: y}, planet.Obstacles}
// 	obstacleCoordinate := coordinate.New(x, y)

// 	didCollide := planetMap.CheckCollision(obstacleCoordinate)

// 	assert.True(t, didCollide)
// }
