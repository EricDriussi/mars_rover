package rover

import (
	"mars_rover/internal/domain/direction"
	"mars_rover/internal/domain/location"
	"mars_rover/internal/domain/planet"
	planetMap "mars_rover/internal/domain/planet_map"
)

type Rover struct {
	location    location.Location
	orientation direction.Direction
	planetMap   planetMap.Map
}

// TODO: should this error if landing on obstacle?
func Land(location location.Location, direction direction.Direction, planet planet.Planet) *Rover {
	return &Rover{location: location, orientation: direction, planetMap: *planetMap.Of(planet)}
}

func (this Rover) Direction() direction.Direction {
	return this.orientation
}

func (this Rover) Location() location.Location {
	return this.location
}

func (this *Rover) MoveForward() {
	ahead := this.location.WillBeAt(this.orientation.RelativePositionAhead(), this.planetMap.Size())
	if this.willHitSomething(ahead) {
		// TODO: how do I "report the obstacle"?
		return
	}
	this.updateLocation(ahead)
}

func (this *Rover) MoveBackward() {
	ahead := this.location.WillBeAt(this.orientation.RelativePositionBehind(), this.planetMap.Size())
	if this.willHitSomething(ahead) {
		// TODO: how do I "report the obstacle"?
		return
	}
	this.updateLocation(ahead)
}

func (this *Rover) TurnLeft() {
	this.orientation = this.orientation.DirectionOnTheLeft()
}

func (this *Rover) TurnRight() {
	this.orientation = this.orientation.DirectionOnTheRight()
}

func (this Rover) willHitSomething(ahead location.Location) bool {
	return this.planetMap.CheckCollision(ahead)
}

func (this *Rover) updateLocation(location location.Location) {
	this.location = location
}
