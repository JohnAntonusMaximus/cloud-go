package main

// Book defines a book
type Book struct {
	Title  string
	Author string
	ISBN   string
}

var books = map[string]Book{
	"123472347": Book{Title: "HitchHiker's Guide To The Galaxy", Author: "Douglas Adams", ISBN: "123472347"},
	"128348123": Book{Title: "Harry Potter", Author: "J.K. Rowling", ISBN: "128348123"},
}

// AllBooks gets all books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

// GetBook gets a book
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new book
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}
