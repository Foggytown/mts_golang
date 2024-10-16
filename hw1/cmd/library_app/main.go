package main

import (
	"fmt"
	"hw1/internal/library"
	"hw1/internal/library/book"
	"hw1/internal/library/searcher/searcher_map"
	"hw1/internal/library/searcher/searcher_slice"
	"strings"
)

func Generate1(name string) string {
	return name
}

func Generate2(name string) string {
	return strings.ReplaceAll(name, "a", "bbb")
}

func main() {
	books := []book.Book{{Name: "War and Love", Popularity: 0}, {Name: "Interesting math and more", Popularity: 25},
		{Name: "One punch man manga", Popularity: 10000}, {Name: "The cat and the witch", Popularity: 0}}

	// create new lib(map storage)
	libMap := library.CreateLibrary(Generate1, searcher_map.MakeSearcherMap())
	for _, book1 := range books {
		libMap.Add(book1)
	}

	//test new lib
	fmt.Printf("Map storage library:\n")
	book1, ok := libMap.Search("War and Love")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	book1, ok = libMap.Search("One punch man manga")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	libMap.SwapGenerator(Generate2)
	book1, ok = libMap.Search("War and Love")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	book1, ok = libMap.Search("Interesting math and more")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}

	// create new lib(slice storage)
	libSlice := library.CreateLibrary(Generate1, searcher_slice.MakeSearcherSlice())
	for _, book1 := range books {
		libSlice.Add(book1)
	}

	//test new lib
	fmt.Printf("Slice storage library:\n")
	book1, ok = libSlice.Search("War and Love")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	book1, ok = libSlice.Search("One punch man manga")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	libSlice.SwapGenerator(Generate2)
	book1, ok = libSlice.Search("War and Love")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}
	book1, ok = libSlice.Search("Interesting math and more")
	if ok {
		book1.Print()
	} else {
		fmt.Println("Can't find book in library")
		return
	}

}
