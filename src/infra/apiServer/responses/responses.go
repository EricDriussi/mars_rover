package responses

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type HTTPResponseHandler interface {
	SendOk(responseBody any)
	SendBadRequest(msg string)
	SendInternalServerError(msg string)
}

type SimpleHTTPResponseHandler struct {
	w http.ResponseWriter
}

func NewResponseHandler(w http.ResponseWriter) SimpleHTTPResponseHandler {
	return SimpleHTTPResponseHandler{w: w}
}

func (this SimpleHTTPResponseHandler) SendOk(responseBody any) {
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

func (this SimpleHTTPResponseHandler) SendBadRequest(msg string) {
	http.Error(this.w, msg, http.StatusBadRequest)
}

func (this SimpleHTTPResponseHandler) SendInternalServerError(msg string) {
	http.Error(this.w, msg, http.StatusInternalServerError)
}
