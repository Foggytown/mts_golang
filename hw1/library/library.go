package library

import "hw1/library/book"
import "hw1/library/searcher"

type Library interface {
	Search(name string) (book.Book, bool)
	GenerateId(name string) string
	SwapGenerator(f func(string) string)
	Add(book book.Book)
}

type Archive struct {
	Generator func(string) string
	Searcher  searcher.Searcher
}

func (a Archive) GenerateId(name string) string {
	return a.Generator(name)
}

func (a Archive) Search(name string) (book.Book, bool) {
	v, ok := a.Searcher.Search(a.GenerateId(name))
	return v, ok
}

func (a Archive) Add(book book.Book) {
	a.Searcher.Add(a.GenerateId(book.Name), book)
}

func (a Archive) SwapGenerator(f func(string) string) {
	a.Generator = f
	a.Searcher.RegenerateId(f)
}

func CreateLibrary(generator func(string) string, search searcher.Searcher) Library {
	arch := Archive{}
	arch.Generator = generator
	arch.Searcher = search
	return Library(arch)
}
