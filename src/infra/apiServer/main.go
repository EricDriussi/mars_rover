package apiServer

import (
	"encoding/json"
	"fmt"
	"log"
	"mars_rover/src/action/create"
	"mars_rover/src/action/move"
	. "mars_rover/src/domain/coordinate/absoluteCoordinate"
	. "mars_rover/src/domain/obstacle/obstacles"
	. "mars_rover/src/domain/rover"
	. "mars_rover/src/infra/apiServer/dto"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"strconv"
	"sync"
)

var repository *SQLiteRepository

var roversMap = make(map[string]Rover)

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	repository = InitFS()

	apiServer := http.NewServeMux()
	// TODO: add load game endpoint
	apiServer.HandleFunc("/api/randomGame", randomGameHandler)
	apiServer.HandleFunc("/api/moveSequence", moveSequenceHandler)

	fmt.Println("API up and running on http://localhost" + port + "/api/")
	log.Fatal(http.ListenAndServe(port, apiServer))
}

func randomGameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Random Game...")
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
