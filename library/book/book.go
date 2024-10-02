package book

import "fmt"

type Book struct {
	Name       string
	Popularity int
}

func (b Book) Print() {
	fmt.Printf("Name: %s, Popularity: %d\n", b.Name, b.Popularity)
}
