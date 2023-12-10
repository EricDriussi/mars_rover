package dto

type CoordinateDTO struct {
	X int
	Y int
}

type RoverDTO struct {
	Id         string
	Coordinate CoordinateDTO
	Direction  string
}

type ObstacleDTO struct {
	Coordinate []CoordinateDTO
}

type PlanetDTO struct {
	Width     int
	Height    int
	Obstacles []ObstacleDTO
}

type CreateResponseDTO struct {
	Rover  RoverDTO
	Planet PlanetDTO
}

type MovementResponseDTO struct {
	Results []SingleMovementDTO
}

type SingleMovementDTO struct {
	Issue      string
	Coordinate CoordinateDTO
	Direction  string
}
