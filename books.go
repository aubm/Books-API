package main

import (
	"net/http"
	"time"

	"github.com/pborman/uuid"
)

// Book is the representation of a book
type Book struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key"`
	Name          string    `json:"name"`
	Author        string    `json:"author"`
	PublishedDate time.Time `json:"publishedDate"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
}

func createBook(w http.ResponseWriter, r *http.Request) {
}

func updateBook(w http.ResponseWriter, r *http.Request) {
}

func getOneBook(w http.ResponseWriter, r *http.Request) {
}

func deleteOneBook(w http.ResponseWriter, r *http.Request) {
}
