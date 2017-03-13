package main

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (book Book) CategoryByLength() string {
	return "foo"
}
