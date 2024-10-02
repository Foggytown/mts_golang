package searcher

import "hw1/library/book"

type Searcher interface {
	Search(id string) (book.Book, bool)
	Add(id string, book book.Book)
	RegenerateId(f func(string) string)
}