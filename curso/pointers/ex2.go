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
