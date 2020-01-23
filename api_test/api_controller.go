package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

/* Answers the '/' api path */
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

/* Answers the '/greet-me/{name}' api path */
func greetMe(w http.ResponseWriter, r *http.Request) {
	/* Extracts the URL parameter which follows the '/greet-me' path */
	name := strings.TrimPrefix(r.URL.Path, "/greet-me/")

	if len(name) > 0 { /* If a parameter {name} exists */

		/* Makes the first character upper case */
		name = strings.Title(strings.ToLower(name))
		/* Returns the greeting */
		fmt.Fprintf(w, "Greetings "+name)

	} else { /* If a parameter {name} doesn't exist */
		fmt.Fprintf(w, "No name found")
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var b []Book = RetrieveBooks()

	json.NewEncoder(w).Encode(b)
}

/* Decides which function should answer to the given api path */
func handleRequests() {

	http.HandleFunc("/", home)
	http.HandleFunc("/greet-me/", greetMe)
	http.HandleFunc("/books", getBooks)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

/* Main method, calls the handling request method on init */
func main() {

	books = []Book{
		Book{Title: "Test1", Author: "Auth1", Pages: 2},
		Book{Title: "Test2", Author: "Auth2", Pages: 3},
	}

	handleRequests()
}
