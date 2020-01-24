package main

import (
	"errors"
)

/* Book struct */
type Book struct {
	Id     int    `json:"Id"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Pages  int    `json:"Pages"`
}

/* Dummy books data */
var books []Book = []Book{
	Book{Id: 1, Title: "Test1", Author: "Auth1", Pages: 2},
	Book{Id: 2, Title: "Test2", Author: "Auth2", Pages: 3},
	Book{Id: 3, Title: "Test1", Author: "Auth2", Pages: 3},
}

func RetrieveBooks() []Book {
	return books
}

func RetrieveBookById(id int) (Book, error) {
	for _, b := range books {
		if b.Id == id {
			return b, nil
		}
	}
	return Book{}, errors.New("No book with the given ID")
}

func RetrieveBooksByTitle(title string) ([]Book, error) {
	var bs []Book
	for _, b := range books {
		if b.Title == title {
			bs = append(bs, b)
		}
	}
	if len(bs) > 0 {
		return bs, nil
	} else {
		return bs, errors.New("No books found with this Title")
	}
}
