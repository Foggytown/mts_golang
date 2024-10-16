package searcher_map

import (
	"hw1/internal/library/book"
)

type SearcherMap struct {
	data map[string]book.Book
}

func (s *SearcherMap) Search(id string) (book.Book, bool) {
	v, ok := s.data[id]
	return v, ok
}

func (s *SearcherMap) Add(id string, book book.Book) {
	s.data[id] = book
}

func (s SearcherMap) RegenerateId(f func(string) string) {
	new_data := make(map[string]book.Book)
	for _, book := range s.data {
		new_data[f(book.Name)] = book
	}
	s.data = new_data
}

func MakeSearcherMap() *SearcherMap {
	res := new(SearcherMap)
	res.data = make(map[string]book.Book)
	return res
}
