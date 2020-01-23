package main

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
}

func RetrieveBooks() []Book {
	return books
}

func RetrieveBookById(id int) Book {
	return books[0]
}
