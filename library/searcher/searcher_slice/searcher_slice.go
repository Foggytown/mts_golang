package searcher_slice

import "hw1/library/book"

type SearcherSlice struct {
	dict map[string]int
	data []book.Book
}

func (s *SearcherSlice) Add(id string, book book.Book) {
	s.dict[id] = len(s.data)
	s.data = append(s.data, book)
}

func (s *SearcherSlice) Search(id string) (book.Book, bool) {
	int_id, ok := s.dict[id]
	if ok {
		return s.data[int_id], ok
	}
	return book.Book{}, ok
}

func (s SearcherSlice) RegenerateId(f func(string) string) {
	new_dict := make(map[string]int)
	for _, int_id := range s.dict {
		new_dict[f(s.data[int_id].Name)] = int_id
	}
	s.dict = new_dict
}

func MakeSearcherSlice() *SearcherSlice {
	res := new(SearcherSlice)
	res.dict = make(map[string]int)
	res.data = make([]book.Book, 0)
	return res
}
