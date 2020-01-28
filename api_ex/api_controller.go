package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

/* Answers the '/' api path */
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

/* Answers the '/greet-me/{name}' api path */
func greetMe(w http.ResponseWriter, r *http.Request) {
	/* Extracts the URL parameter which follows the '/greet-me/' path */
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

/* Retrieves a list with all the books */
func getBooks(w http.ResponseWriter, r *http.Request) {
	var b []Book = RetrieveBooks()

	json.NewEncoder(w).Encode(b)
}

/* Retrieves a book with the given ID */
func getBookById(w http.ResponseWriter, r *http.Request) {
	/* Extracts the URL parameter which follows the '/books/' path */
	id, e := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/books/"))

	if e != nil {
		fmt.Fprintf(w, "No valid ID")
	} else {
		b, err := RetrieveBookById(id)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			json.NewEncoder(w).Encode(b)
		}
	}
}

func getBookByTitle(w http.ResponseWriter, r *http.Request) {
	title := strings.TrimPrefix(r.URL.Path, "/books/title/")

	bs, err := RetrieveBooksByTitle(title)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	} else {
		json.NewEncoder(w).Encode(bs)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		/* there's an error reading body data*/
	} else {
		fmt.Fprintf(w, "creating\n"+string(body))
		fmt.Println("creating")
	}
}

/* Decides which function should answer to the given api path */
func handleRequests() {

	http.HandleFunc("/", home)
	http.HandleFunc("/greet-me/", greetMe)
	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/book", createBook)
	http.HandleFunc("/book/", getBookById)
	http.HandleFunc("/books/title/", getBookByTitle)

	log.Fatal(http.ListenAndServe(":10000", nil))
}

/* Main method, calls the handling request method on init */
func main() {

	handleRequests()
}
