package action

import (
	"fmt"
	. "mars_rover/src/domain/rover"
)

type MovementResult struct {
	MovedRover Rover
	Collisions *Collisions
}

type Collisions struct {
	collisionList []Collision
}

func (this *Collisions) Add(command string, err error) {
	this.collisionList = append(this.collisionList, Collision{command: command, err: err})
}

func (this *Collisions) List() []Collision {
	return this.collisionList
}

func (this *Collisions) AsStringArray() []string {
	var collisions []string
	for _, collision := range this.collisionList {
		collisions = append(collisions, collision.AsString())
	}
	return collisions
}

type Collision struct {
	command string
	err     error
}

func (this *Collision) AsString() string {
	return fmt.Sprintf("error executing command %v: %v", this.command, this.err)
}
