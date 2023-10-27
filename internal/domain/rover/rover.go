package rover

import (
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/planet_map"
	"mars_rover/internal/domain/position"
)

type Rover struct {
	position  position.Position
	direction direction.Direction
	planetMap planetMap.PlanetMap
}

func Land(position position.Position, direction direction.Direction, planet planet.Planet) *Rover {
	return &Rover{position: position, direction: direction, planetMap: *planetMap.Of(planet)}
}

func (this Rover) Direction() direction.Direction {
	return this.direction
}

func (this Rover) Position() position.Position {
	return this.position
}

func (this *Rover) MoveForward() {
	this.position.AddOrWrap(this.direction.RelativePositionAhead(), this.planetMap.Size())
}

func (this *Rover) MoveBackward() {
	this.position.AddOrWrap(this.direction.RelativePositionBehind(), this.planetMap.Size())
}

func (this *Rover) TurnLeft() {
	this.direction = this.direction.DirectionOnTheLeft()
}

func (this *Rover) TurnRight() {
	this.direction = this.direction.DirectionOnTheRight()
}

func (this Rover) CheckObstacle() bool {
	return this.planetMap.CheckCollision(this.position)
}
