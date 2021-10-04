package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	_"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestGetDocuments(t *testing.T) {
	router := gin.Default()
	router.GET("/documents" , GetDocuments )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expected := `[{"id":"1","nom":"Nom_1","description":"description"},{"id":"2","nom":"Nom_2","description":"description"},{"id":"3","nom":"Nom_3","description":"description"},{"id":"4","nom":"Nom_4","description":"description"},{"id":"5","nom":"Nom_5","description":"description"},{"id":"6","nom":"Nom_6","description":"description"}]`
	assert.Equal(t, expected, w.Body.String())	
}

func TestGetDocumentByID(t *testing.T) {
	router := gin.Default()
	router.GET("/documents/:id" , GetDocumentByID )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expected := `{"id":"1","nom":"Nom_1","description":"description"}`
	assert.Equal(t, expected, w.Body.String())	
} 


func TestGetDocumentByIDNotFound(t *testing.T) {
	router := gin.Default()
	router.GET("/documents/:id" , GetDocumentByID )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/documents/1000", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}


func TestNewDocument(t *testing.T) {
	documentsLength := len(documents)
	var document = []byte(`{"ID": "7", "Nom": "Nom_7", "Description": "description"}`)
	router := gin.Default()
	router.POST("/documents" , NewDocument )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/documents", bytes.NewBuffer(document))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t , len(documents) , documentsLength+1 )
}  



func TestRemoveDocument(t *testing.T) {
	router := gin.Default()
	router.DELETE("/documents/:id" , GetDocumentByID )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/documents/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	expected := `{"id":"1","nom":"Nom_1","description":"description"}`
	assert.Equal(t, expected, w.Body.String())
}


func TestRemoveNonExistDocument(t *testing.T) {
	router := gin.Default()
	router.DELETE("/documents/:id" , GetDocumentByID )
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/documents/1000", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
