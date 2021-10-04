package main

import (
	_"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)


//la structure document
type document struct {
    ID     string  `json:"id"`
    Nom  string  `json:"nom"`
    Description string  `json:"description"`
}

/**
Données en mémoire , habituellement , stockées sur un RDMS : mysql, pgsql, ...
in-memorydata base /NoSQL : SQLlite - Redis - MongoDB, ... 
**/
var documents []document


//Initialisation des données 
func init() {
	documents = append(documents, document {ID: "1", Nom: "Nom_1", Description: "description"} )
	documents = append(documents, document {ID: "2", Nom: "Nom_2", Description: "description"} )
	documents = append(documents, document {ID: "3", Nom: "Nom_3", Description: "description"} )
	documents = append(documents, document {ID: "4", Nom: "Nom_4", Description: "description"} )
	documents = append(documents, document {ID: "5", Nom: "Nom_5", Description: "description"} )
	documents = append(documents, document {ID: "6", Nom: "Nom_6", Description: "description"} )
}


//Retourne la list de tout les documents.
func GetDocuments(c *gin.Context) {
    c.JSON(http.StatusOK, documents )
}

//Retourne un nouveau document par ID 
func GetDocumentByID(c *gin.Context) { 
    idDocument := c.Param("id") 
    for key, val := range documents {
        if val.ID == idDocument {
            c.JSON(http.StatusOK, documents[key])
            return 
        }
    }
    message := gin.H{"message": "document with id#" + idDocument + " not found"}
    c.JSON(http.StatusNotFound, message ) 
} 

//Ajout d'un nouveau document
func NewDocument(c *gin.Context) {
    var doc document
    if err := c.BindJSON(&doc); err != nil {
        return
    }
    documents = append(documents, doc )
    c.JSON(http.StatusCreated, documents ) 
}

//Suppression d'un document 
func RemoveDocument(c *gin.Context) {
    idDocument := c.Param("id")
    for key, val := range documents {
        if val.ID == idDocument {
            documents = append(documents[:key], documents[key+1:]...)
            c.JSON(http.StatusOK, gin.H{"message": "the document with id#" + idDocument +" was removed"} )
            return
        }
    } 
    message := gin.H{"message": "no such document with id#" + idDocument }
    c.JSON(http.StatusNotFound, message )  
}


//main
func main() {
    router := gin.Default() 
	router.GET("/documents", GetDocuments)
    router.GET("/documents/:id", GetDocumentByID)
	router.POST("/documents",  NewDocument)
	router.DELETE("/documents/:id", RemoveDocument) 
	router.Run(":8080")
}