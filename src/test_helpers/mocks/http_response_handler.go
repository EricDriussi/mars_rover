package mocks

import . "github.com/stretchr/testify/mock"

type MockHTTPResponseHandler struct {
	Mock
}

func (this *MockHTTPResponseHandler) SendOk(responseBody any) {
	this.Called(responseBody)
}

func (this *MockHTTPResponseHandler) SendBadRequest(msg string) {
	this.Called(msg)
}

func (this *MockHTTPResponseHandler) SendInternalServerError(msg string) {
	this.Called(msg)
}
