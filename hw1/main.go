package main

import (
	"fmt"
	"hw1/library"
	"hw1/library/book"
	"hw1/library/searcher/searcher_map"
	"hw1/library/searcher/searcher_slice"
	"strings"
)

func Generate1(name string) string {
	return name
}

func Generate2(name string) string {
	return strings.ReplaceAll(name, "a", "bbb")
}

func main() {
	books := []book.Book{book.Book{Name: "War and Love", Popularity: 0}, book.Book{Name: "Intresting math and more", Popularity: 25},
		book.Book{Name: "One punch man manga", Popularity: 10000}, book.Book{Name: "The cat and the witch", Popularity: 0}}

	// create new lib(map storage)
	lib_map := library.CreateLibrary(Generate1, searcher_map.MakeSearcherMap())
	for _, book1 := range books {
		lib_map.Add(book1)
	}

	//test new lib
	fmt.Printf("Map storage library:\n")
	book1, ok := lib_map.Search("War and Love")
	if ok {
		book1.Print()
	}
	book1, ok = lib_map.Search("One punch man manga")
	if ok {
		book1.Print()
	}
	lib_map.SwapGenerator(Generate2)
	book1, ok = lib_map.Search("War and Love")
	if ok {
		book1.Print()
	}
	book1, ok = lib_map.Search("Intresting math and more")
	if ok {
		book1.Print()
	}

	// create new lib(slice storage)
	lib_slice := library.CreateLibrary(Generate1, searcher_slice.MakeSearcherSlice())
	for _, book1 := range books {
		lib_slice.Add(book1)
	}

	//test new lib
	fmt.Printf("Slice storage library:\n")
	book1, ok = lib_slice.Search("War and Love")
	if ok {
		book1.Print()
	}
	book1, ok = lib_slice.Search("One punch man manga")
	if ok {
		book1.Print()
	}
	lib_slice.SwapGenerator(Generate2)
	book1, ok = lib_slice.Search("War and Love")
	if ok {
		book1.Print()
	}
	book1, ok = lib_slice.Search("Intresting math and more")
	if ok {
		book1.Print()
	}

}
