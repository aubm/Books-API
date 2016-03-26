package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

// Book is the representation of a book
type Book struct {
	ID              string    `json:"id" sql:"type:varchar(36);primary key" gorm:"primary_key"`
	Name            string    `json:"name"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publicationDate"`
	LibraryID       string    `json:"-"`
	Library         Library   `json:"-"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	if libraryExists := checkLibrary(w, r); libraryExists == false {
		return
	}

	var books []Book
	db.Where("library_id = ?", mux.Vars(r)["libraryId"]).Find(&books)
	writeJSON(w, books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	if libraryExists := checkLibrary(w, r); libraryExists == false {
		return
	}

	newBook, err := bookFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON provided or invalid date time format", http.StatusBadRequest)
		return
	}
	validationErrs := validateBook(*newBook)
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, validationErrs)
		return
	}
	newBook.ID = uuid.NewUUID().String()
	newBook.LibraryID = mux.Vars(r)["libraryId"]
	db.Create(newBook)
	w.Header().Set("Book-Id", newBook.ID)
	w.WriteHeader(201)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	if libraryExists := checkLibrary(w, r); libraryExists == false {
		return
	}

	book, err := findOneBook(mux.Vars(r)["bookId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	update, err := bookFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON provided or invalid date time format", http.StatusBadRequest)
		return
	}
	book.Name = update.Name
	book.Author = update.Author
	book.PublicationDate = update.PublicationDate

	validationErrs := validateBook(book)
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, validationErrs)
		return
	}

	writeJSON(w, book)
}

func getOneBook(w http.ResponseWriter, r *http.Request) {
	if libraryExists := checkLibrary(w, r); libraryExists == false {
		return
	}

	book, err := findOneBook(mux.Vars(r)["bookId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	writeJSON(w, book)
}

func deleteOneBook(w http.ResponseWriter, r *http.Request) {
	if libraryExists := checkLibrary(w, r); libraryExists == false {
		return
	}

	book, err := findOneBook(mux.Vars(r)["bookId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	db.Delete(&book)
	w.WriteHeader(http.StatusNoContent)
}

func checkLibrary(w http.ResponseWriter, r *http.Request) bool {
	_, err := findOneLibrary(mux.Vars(r)["libraryId"])
	if err != nil {
		http.Error(w, "Library not found", http.StatusNotFound)
		return false
	}
	return true
}

func bookFromJSON(data io.ReadCloser) (*Book, error) {
	var book = new(Book)
	decoder := json.NewDecoder(data)
	err := decoder.Decode(book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func validateBook(book Book) (errs []string) {
	if book.Name == "" {
		errs = append(errs, "The name is required")
	}
	if book.Author == "" {
		errs = append(errs, "The author is required")
	}
	return
}

func findOneBook(bookID string) (Book, error) {
	var book Book
	db.Where("id = ?", bookID).Find(&book)
	if book.ID == "" {
		return Book{}, errors.New("Book not found")
	}
	return book, nil
}
