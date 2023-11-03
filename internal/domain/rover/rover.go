package rover

import (
	"mars_rover/internal/domain/coordinate"
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

func (this *Rover) MoveForward() {
	ahead := this.location.AheadWillBeAt(this.planetMap.Size())
	if this.willHitSomething(ahead) {
		// TODO: how do I "report the obstacle"?
		return
	}
	this.location.UpdateCoordinate(ahead)
}

func (this *Rover) MoveBackward() {
	behind := this.location.BehindWillBeAt(this.planetMap.Size())
	if this.willHitSomething(behind) {
		// TODO: how do I "report the obstacle"?
		return
	}
	this.location.UpdateCoordinate(behind)
}

func (this *Rover) TurnLeft() {
	this.location.UpdateWithDirectionOnTheLeft()
}

func (this *Rover) TurnRight() {
	this.location.UpdateWithDirectionOnTheRight()
}

func (this Rover) willHitSomething(futureCoord coordinate.AbsoluteCoordinate) bool {
	return this.planetMap.CheckCollision(futureCoord)
}
