package apiServer

import (
	"encoding/json"
	"fmt"
	"log"
	. "mars_rover/src/action"
	"mars_rover/src/action/createRandom/bounded"
	"mars_rover/src/action/load"
	"mars_rover/src/action/move/resilient"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/apiServer/responses"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"sync"
)

var creationAction CreateRandomAction
var movementAction MoveAction
var loadAction LoadAction

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	creationAction = boundedRandomCreator.With(InitFS())
	movementAction = resilient_mover.With(InitFS())
	loadAction = load.With(InitFS())

	apiServer := http.NewServeMux()
	apiServer.HandleFunc("/api/randomGame", corsMiddleWare(randomGameHandler))
	apiServer.HandleFunc("/api/moveSequence", corsMiddleWare(moveSequenceHandler))
	apiServer.HandleFunc("/api/loadGame", corsMiddleWare(loadHandler))

	fmt.Println("API up and running on http://localhost" + port + "/api/")
	log.Fatal(http.ListenAndServe(port, apiServer))
}

func randomGameHandler(w http.ResponseWriter, r *http.Request) {
	responseHandler := NewResponseHandler(w)
	if r.Method != "POST" {
		responseHandler.SendBadRequest("Invalid method")
		return
	}

	RandomGame(creationAction, responseHandler)
}

func moveSequenceHandler(w http.ResponseWriter, r *http.Request) {
	responseHandler := NewResponseHandler(w)
	if r.Method != "POST" {
		responseHandler.SendBadRequest("Invalid method")
		return
	}

	var request MoveRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseHandler.SendBadRequest("Invalid request body")
		return
	}

	MoveRover(movementAction, request, responseHandler)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	responseHandler := NewResponseHandler(w)
	if r.Method != "POST" {
		responseHandler.SendBadRequest("Invalid method")
		return
	}

	var request LoadRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseHandler.SendBadRequest("Invalid request body")
		return
	}

	LoadGame(loadAction, request, responseHandler)
}

func corsMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:6969")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		// TODO.LM: I don't know why, but some browsers send an OPTIONS request before the actual request to check for cors ¯\_(ツ)_/¯
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}
