package apiServer

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mars_rover/src/action"
	. "mars_rover/src/action"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/persistence"
	"net/http"
	"strconv"
	"sync"
)

var act *LaxAction

func RunOn(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	// TODO.LM: Only one action being used by multiple controllers: is that OK?
	act = action.For(InitFS())

	apiServer := http.NewServeMux()
	// TODO: add load game endpoint
	apiServer.HandleFunc("/api/randomGame", randomGameHandler)
	apiServer.HandleFunc("/api/moveSequence", moveSequenceHandler)

	fmt.Println("API up and running on http://localhost" + port + "/api/")
	log.Fatal(http.ListenAndServe(port, apiServer))
}

func randomGameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		sendBadRequest(w, "Invalid method")
		return
	}

	responseDTO, err := RandomGame(act)
	if err != nil {
		sendInternalServerError(w, err.Error())
		return
	}
	sendOkResponse(w, responseDTO)
}

func moveSequenceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		sendBadRequest(w, "Invalid method")
		return
	}

	var request MoveRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		sendBadRequest(w, "Invalid request body")
		return
	}
	_, err = uuid.Parse(request.Id)
	if err != nil {
		sendBadRequest(w, "Invalid ID")
		return
	}

	responseDTO, err := MoveRover(act, request)
	if err != nil {
		sendInternalServerError(w, err.Error())
		return
	}
	sendOkResponse(w, responseDTO)
}

func sendOkResponse(w http.ResponseWriter, responseDTO any) {
	jsonResponse, err := json.Marshal(responseDTO)
	if err != nil {
		sendInternalServerError(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(jsonResponse)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		sendInternalServerError(w, err.Error())
		return
	}
}

func sendBadRequest(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}

func sendInternalServerError(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}
