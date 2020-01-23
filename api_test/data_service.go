package main

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Pages  int    `json:"Pages"`
}

var books []Book

func RetrieveBooks() []Book {
	return books
}
