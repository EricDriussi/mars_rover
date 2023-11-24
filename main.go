package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	. "mars_rover/internal/domain/coordinate/absoluteCoordinate"
	. "mars_rover/internal/domain/obstacle/obstacles"
	. "mars_rover/internal/infra"
	"mars_rover/internal/use_case/create"
	"mars_rover/internal/use_case/move"
	"net/http"
	"strconv"
)

var (
	repository *SQLiteRepository
	db         *sql.DB
)

const PORT = ":8080"

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

func main() {
	http.HandleFunc("/api/randomRover", randomRoverHandler)
	http.HandleFunc("/api/moveSequence", moveSequenceHandler)

	fmt.Println("Ready to go!")
	fmt.Println("Listening on http://localhost" + PORT + "/api/")
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func randomRoverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Random Value...")
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
	fmt.Println("Random Value Created!")
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
	fmt.Println("Moving Value...")
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

	moveAction, err := move.For2(request.Id, repository)
	errs := moveAction.MoveSequence(request.Commands)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(errs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("Done Moving Value!")
}
