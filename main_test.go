package main

import (
	
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPlanetor(t *testing.T){

	mockResponse := `{"message":"DB filled with dummydata"}`
	
	testRouter := SetupRouter()
	
    req, _ := http.NewRequest("GET", "/seed", nil)
    w := httptest.NewRecorder()
    testRouter.ServeHTTP(w, req)

    responseData, _ := ioutil.ReadAll(w.Body)
    assert.Equal(t, mockResponse, string(responseData))
    assert.Equal(t, http.StatusOK, w.Code)
	
	
}

