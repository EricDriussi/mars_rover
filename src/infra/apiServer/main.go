package apiServer

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mars_rover/src/action/create"
	"mars_rover/src/action/move"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"strconv"
	"sync"
)

var repository *SQLiteRepository

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
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	action := create.For(repository)
	jsonResponse, err := RandomGame(*action)
	sendResponse(w, jsonResponse, err)
}

func moveSequenceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var request MoveRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	_, err = uuid.Parse(request.Id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	moveAction := move.For(repository)
	jsonResponse, err := MoveRover(*moveAction, request)
	sendResponse(w, jsonResponse, err)
}

func sendResponse(w http.ResponseWriter, jsonResponse []byte, err error) {
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
