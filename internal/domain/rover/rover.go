package rover

import (
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/rover/planet_map"
)

type Rover struct {
	location  location.Location
	planetMap planetMap.Map
}

// TODO: should this error if landing on obstacle?
func Land(location location.Location, planet planet.Planet) *Rover {
	return &Rover{location: location, planetMap: *planetMap.Of(planet)}
}

func (this Rover) Location() *location.Location {
	return &this.location
}

func (this *Rover) TurnLeft() {
	this.location.FaceLeft()
}

func (this *Rover) TurnRight() {
	this.location.FaceRight()
}

// TODO: add tests with mocks
func (this *Rover) MoveForward() {
	this.location.CalculatePositionAhead()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		// TODO: how do I "report the obstacle"?
		this.location.Reset()
		return
	}
	this.location.UpdatePosition()
}

// TODO: add tests with mocks
func (this *Rover) MoveBackward() {
	this.location.CalculatePositionBehind()
	if this.willBeOutOfBounds() {
		this.location.WrapAround(this.planetMap.Size())
	}
	if this.willHitSomething() {
		// TODO: how do I "report the obstacle"?
		this.location.Reset()
		return
	}
	this.location.UpdatePosition()
}

func (this Rover) willBeOutOfBounds() bool {
	return this.planetMap.IsOutOfBounds(this.location.WillBeAt())
}

func (this Rover) willHitSomething() bool {
	return this.planetMap.CollidesWithObstacle(this.location.WillBeAt())
}
