package apiServer

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
	Rover  RoverDTO
	Errors []string
}
