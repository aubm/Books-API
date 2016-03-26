package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var env = flag.String("env", "dev", "the execution environnement, use prod in production")

func main() {
	flag.Parse()

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
	apiRouter.HandleFunc("/libraries/{libraryId}/books", createBook).Methods("POST")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", updateBook).Methods("PUT")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", getOneBook).Methods("GET")
	apiRouter.HandleFunc("/libraries/{libraryId}/books/{bookId}", deleteOneBook).Methods("DELETE")

	n := negroni.New()
	if *env == "dev" {
		n.Use(resetDBMiddleware{})
	}
	n.Use(negroni.Wrap(router))

	n.Run(":8080")
}

var initContext sync.Once

func initDatabase() {
	initContext.Do(func() {
		dbName := os.Getenv("MYSQL_DATABASE")
		dbUser := os.Getenv("MYSQL_USER")
		dbPassword := os.Getenv("MYSQL_PASSWORD")
		log.Printf("database config, dbname: %v, dbuser: %v, dbpassword: %v", dbName, dbUser, dbPassword)
		initConnection := func() error {
			var err error
			db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(db:3306)/%v?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbName))
			if err != nil {
				return err
			}

			_, err = db.Raw("SELECT 1 FROM libraries").Rows()
			if err != nil {
				log.Print("database is not ready, executing SQL scripts reset.sql and data.sql")
				executeSQLFile("reset")
				executeSQLFile("data")
			}

			return nil
		}
		log.Print("attempt to init database connection")
		for {
			err := initConnection()
			if err == nil {
				log.Print("connection success")
				return
			} else {
				log.Printf("connection failed, will retry in 5 seconds, reason: %v", err)
				time.Sleep(time.Second * 5)
			}
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

type resetDBMiddleware struct{}

func (mw resetDBMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	dbScripts := r.Header.Get("dbscripts")
	if dbScripts != "" {
		scripts := strings.Split(dbScripts, "|")
		for _, script := range scripts {
			log.Printf("executing script %v.sql", script)
			executeSQLFile(script)
		}
	}
	next(rw, r)
}

func executeSQLFile(filename string) {
	contents, err := ioutil.ReadFile("./sql_scripts/" + filename + ".sql")
	if err != nil {
		log.Printf("Failed to open script file: %v", err)
		return
	}

	queries := strings.Split(string(contents), ";")
	for _, query := range queries {
		query = strings.Trim(query, " ")
		if query != "" {
			db.Exec(query)
		}
	}
}
