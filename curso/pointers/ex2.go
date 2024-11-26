package main

import "fmt"

/*
	Define two structs, PrintedBook and AudioBook. Both should implement the interface Book, which has the following methods:

	Title() string
	Author() string
	Content() string

	Both structs should have the following fields:
	PrintedBook: title, author, pageCount.
	AudioBook: title, author, lengthInMinutes.

	Implement the methods for both structs.

*/

type Book interface {
	Title() string
	Author() string
	Content()
}

type PrintedBook struct {
	title     string
	author    string
	pageCount int
}

type AudioBook struct {
	title           string
	author          string
	lengthInMinutes int
}

func (b PrintedBook) Title() string {
	return fmt.Sprintf(b.title)
}

func (b AudioBook) Title() string {
	return fmt.Sprintf(b.title)
}

func (b PrintedBook) Author() string {
	return fmt.Sprintf(b.author)
}

func (b AudioBook) Author() string {
	return fmt.Sprintf(b.author)
}

func (b PrintedBook) Content() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Author: ", b.author)
	fmt.Println("Time: ", b.pageCount)
}

func (b AudioBook) Content() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Author: ", b.author)
	fmt.Println("Time: ", b.lengthInMinutes)
}

func (b *PrintedBook) setAuthor(a string) {
	b.author = a
}

func ex3() {
	c := AudioBook{
		title:           "El extranjero",
		author:          "Albert Cammus",
		lengthInMinutes: 60,
	}

	b := PrintedBook{
		title:     "100 años de soledad",
		author:    "Gabriel García Márquez",
		pageCount: 850,
	}
	c.Content()
	b.Content()

	author := b.Author()
	fmt.Println(author)
	b.setAuthor("El Gabo")
	author2 := b.author
	fmt.Println(author2)
	b.Content()

}
