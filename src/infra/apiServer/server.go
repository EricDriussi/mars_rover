package apiServer

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"mars_rover/src/action"
	. "mars_rover/src/action"
	. "mars_rover/src/infra/apiServer/controllers"
	. "mars_rover/src/infra/apiServer/responses"
	. "mars_rover/src/infra/persistence"
	"net/http"
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
	responseHandler := NewResponseHandler(w)
	if r.Method != "POST" {
		responseHandler.SendBadRequest("Invalid method")
		return
	}

	RandomGame(act, responseHandler)
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
	_, err = uuid.Parse(request.Id)
	if err != nil {
		responseHandler.SendBadRequest("Invalid ID")
		return
	}

	MoveRover(act, request, responseHandler)
}
