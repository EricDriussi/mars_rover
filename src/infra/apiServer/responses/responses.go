package responses

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ResponseHandler struct {
	w http.ResponseWriter
}

func NewResponseHandler(w http.ResponseWriter) ResponseHandler {
	return ResponseHandler{w: w}
}

func (this ResponseHandler) SendOk(responseBody any) {
	jsonResponse, err := json.Marshal(responseBody)
	if err != nil {
		this.SendInternalServerError(err.Error())
		return
	}

	this.w.Header().Set("Content-Type", "application/json")
	this.w.Header().Set("Content-Length", strconv.Itoa(len(jsonResponse)))
	this.w.WriteHeader(http.StatusOK)
	_, err = this.w.Write(jsonResponse)
	if err != nil {
		this.SendInternalServerError(err.Error())
		return
	}
}

func (this ResponseHandler) SendBadRequest(msg string) {
	http.Error(this.w, msg, http.StatusBadRequest)
}

func (this ResponseHandler) SendInternalServerError(msg string) {
	http.Error(this.w, msg, http.StatusInternalServerError)
}
