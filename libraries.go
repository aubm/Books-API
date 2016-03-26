package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pborman/uuid"
)

// Library is the representation of a library
type Library struct {
	ID      string `json:"id" sql:"type:varchar(36);primary key" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Books   []Book `json:"-"`
}

func getLibraries(w http.ResponseWriter, r *http.Request) {
	var libraries []Library
	db.Find(&libraries)
	writeJSON(w, libraries)
}

func createLibrary(w http.ResponseWriter, r *http.Request) {
	newLibrary, err := libraryFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON provided", http.StatusBadRequest)
		return
	}
	validationErrs := validateLibrary(*newLibrary)
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, validationErrs)
		return
	}
	newLibrary.ID = uuid.NewUUID().String()
	db.Create(newLibrary)
	w.Header().Set("Library-Id", newLibrary.ID)
	w.WriteHeader(201)
}

func updateLibrary(w http.ResponseWriter, r *http.Request) {
	library, err := findOneLibrary(mux.Vars(r)["libraryId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	update, err := libraryFromJSON(r.Body)
	if err != nil {
		http.Error(w, "Invalid JSON provided", http.StatusBadRequest)
		return
	}
	library.Name = update.Name
	library.Address = update.Address
	library.Phone = update.Phone

	validationErrs := validateLibrary(library)
	if len(validationErrs) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(w, validationErrs)
		return
	}

	writeJSON(w, library)
}

func getOneLibrary(w http.ResponseWriter, r *http.Request) {
	library, err := findOneLibrary(mux.Vars(r)["libraryId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	writeJSON(w, library)
}

func deleteLibrary(w http.ResponseWriter, r *http.Request) {
	library, err := findOneLibrary(mux.Vars(r)["libraryId"])
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	db.Where("library_id = ?", library.ID).Delete(&Book{})
	db.Delete(&library)
	w.WriteHeader(http.StatusNoContent)
}

func libraryFromJSON(data io.ReadCloser) (*Library, error) {
	var library = new(Library)
	decoder := json.NewDecoder(data)
	err := decoder.Decode(library)
	if err != nil {
		return nil, err
	}
	return library, nil
}

func validateLibrary(library Library) (errs []string) {
	if library.Name == "" {
		errs = append(errs, "The name is required")
	}
	if library.Address == "" {
		errs = append(errs, "The address is required")
	}
	if library.Phone == "" {
		errs = append(errs, "The phone is required")
	}
	return
}

func findOneLibrary(libraryID string) (Library, error) {
	var library Library
	db.Where("id = ?", libraryID).Find(&library)
	if library.ID == "" {
		return Library{}, errors.New("Library not found")
	}
	return library, nil
}
