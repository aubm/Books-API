package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	initDatabase()
	defer db.Close()

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/libraries", getLibraries).Methods("GET")
	apiRouter.HandleFunc("/libraries", createLibrary).Methods("POST")
	apiRouter.HandleFunc("/libraries/{libraryId}", updateLibrary).Methods("PUT")
	apiRouter.HandleFunc("/libraries/{libraryId}", getOneLibrary).Methods("GET")
	apiRouter.HandleFunc("/libraries/{libraryId}", deleteLibrary).Methods("DELETE")
	apiRouter.HandleFunc("/libraries/{libraryId}/books", getBooks).Methods("GET")
	apiRouter.HandleFunc("/libraries/{libraryId}/books", createBook).Methods("GET")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", updateBook).Methods("PUT")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", getOneBook).Methods("GET")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", deleteOneBook).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var initContext sync.Once

func initDatabase() {
	initContext.Do(func() {
		var err error
		db, err = gorm.Open("mysql", "root:root@/books_api?charset=utf8&parseTime=True&loc=Local")
		if err != nil {
			log.Fatalf("Fail to connect to db: %v", err)
		}
	})
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("error while encoding json: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
