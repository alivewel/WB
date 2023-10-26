package main

import (
	"02/multiset"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

func main() {
	ms := multiset.New[string]()
	ms.Insert("cat")
	ms.Insert("cat")
	ms.Insert("dog")
	ms.Insert("cat")
	ms.Insert("tree")
	ms.Print()
}
