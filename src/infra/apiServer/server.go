package apiServer

import (
	"encoding/json"
	"fmt"
	"log"
	. "mars_rover/src/action"
	boundedRandomGameCreator "mars_rover/src/action/createRandom/bounded"
	"mars_rover/src/action/gameLoader"
	resilientMover "mars_rover/src/action/move/resilient"
	. "mars_rover/src/infra/apiServer/controllers"
	"mars_rover/src/infra/apiServer/responses"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"sync"
)

var createAction CreateRandomAction
var moveAction MoveAction
var loadAction LoadAction

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	sqliteRepo := InitFS()
	createAction = boundedRandomGameCreator.With(sqliteRepo)
	moveAction = resilientMover.With(sqliteRepo)
	loadAction = gameLoader.With(sqliteRepo)

	apiServer := http.NewServeMux()
	setupPaths(apiServer)
	fmt.Println("API up and running on http://localhost" + port + "/api/")
	log.Fatal(http.ListenAndServe(port, apiServer))
}

func setupPaths(apiServer *http.ServeMux) {
	apiServer.HandleFunc("/api/randomGame", corsMiddleWare(onlyPostVerbMiddleWare(randomGameHandler)))
	apiServer.HandleFunc("/api/moveSequence", corsMiddleWare(onlyPostVerbMiddleWare(moveSequenceHandler)))
	apiServer.HandleFunc("/api/loadGame", corsMiddleWare(onlyPostVerbMiddleWare(loadHandler)))
}

func randomGameHandler(w http.ResponseWriter, _ *http.Request) {
	RandomGame(createAction, responses.NewResponseHandler(w))
}

func moveSequenceHandler(w http.ResponseWriter, r *http.Request) {
	responseHandler := responses.NewResponseHandler(w)
	var request MoveRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		responseHandler.SendBadRequest("Invalid request body")
		return
	}

	MoveRover(moveAction, request, responseHandler)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	responseHandler := responses.NewResponseHandler(w)
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
			responseHandler := responses.NewResponseHandler(w)
			responseHandler.SendOk("CORS check")
			return
		}
		next.ServeHTTP(w, r)
	}
}

func onlyPostVerbMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			responseHandler := responses.NewResponseHandler(w)
			responseHandler.SendBadRequest("Invalid HTTP verb")
			return
		}
		next.ServeHTTP(w, r)
	}
}
