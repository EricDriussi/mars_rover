package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mars_rover/src/action/create"
	"mars_rover/src/action/move"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"strconv"
)

var (
	repository *SQLiteRepository
	db         *sql.DB
)

const PORT = ":8080"

var roversMap = make(map[string]Rover)

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

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/api/randomGame", randomGameHandler)
	http.HandleFunc("/api/moveSequence", moveSequenceHandler)

	fmt.Println("Ready to go!")
	fmt.Println("Listening on http://localhost" + PORT + "/api/")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func randomGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Random Rover...")
	db, repository = InitFS()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(db)

	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}
	curiosity, err := create.Random(repository)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	roversMap[curiosity.Id().String()] = curiosity
	coordinate := curiosity.Coordinate()
	m := curiosity.Map()
	response := CreateResponseDTO{
		Rover: RoverDTO{
			Id: curiosity.Id().String(),
			Coordinate: CoordinateDTO{
				X: coordinate.X(),
				Y: coordinate.Y(),
			},
			Direction: curiosity.Direction().CardinalPoint(),
		},
		Planet: PlanetDTO{
			Width:     m.Width(),
			Height:    m.Height(),
			Obstacles: mapDomainToDTOObstacles(m.Obstacles()),
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonResponse)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Random Rover Created!")
}

func mapDomainToDTOObstacles(obstacles Obstacles) []ObstacleDTO {
	var obstaclesDTO []ObstacleDTO
	for _, obstacle := range obstacles.List() {
		coordinates := obstacle.Coordinates()
		obstaclesDTO = append(obstaclesDTO, ObstacleDTO{
			Coordinate: mapDomainToDTOCoordinates(coordinates),
		})
	}
	return obstaclesDTO
}

func mapDomainToDTOCoordinates(c []AbsoluteCoordinate) []CoordinateDTO {
	var coordinatesDTO []CoordinateDTO
	for _, coordinate := range c {
		coordinatesDTO = append(coordinatesDTO, CoordinateDTO{
			X: coordinate.X(),
			Y: coordinate.Y(),
		})
	}
	return coordinatesDTO
}

func moveSequenceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Moving Rover...")
	db, repository = InitFS()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(db)

	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}
	var request struct {
		Commands string `json:"commands"`
		Id       string `json:"id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	rover, exists := roversMap[request.Id]
	if !exists {
		http.Error(w, "Rover not found", http.StatusBadRequest)
		return
	}

	moveAction := move.For(repository)
	movementResult := moveAction.MoveSequence(rover, request.Commands)
	if movementResult.Error != nil {
		http.Error(w, fmt.Sprintf("Unexpected error, aborting: %v", movementResult.Error.Error()), http.StatusInternalServerError)
		return
	}

	updatedRover := movementResult.Rover
	coordinate := updatedRover.Coordinate()
	roverToReturn := RoverDTO{
		Id: updatedRover.Id().String(),
		Coordinate: CoordinateDTO{
			X: coordinate.X(),
			Y: coordinate.Y(),
		},
		Direction: updatedRover.Direction().CardinalPoint(),
	}

	response := MovementResponseDTO{
		// TODO: returning the rover is not enough, should return a list of coordinates-directions since one command might fail but the rover can keep moving
		// Decide in front end if paint all positions or just the last one
		Rover: roverToReturn,
		// TODO: these are not "Errors", they are collisions
		Errors: movementResult.MovementErrors.AsStringArray(),
		// TODO: what about non-movement errors?
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonResponse)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
