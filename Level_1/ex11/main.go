package main

import "fmt"

func intersec(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for elem := range set1 {
		// проверка, существует ли element во втором множестве set2
		if _, exists := set2[elem]; exists {
			result[elem] = struct{}{}
		}
	}

	return result
}

// менее оптимизированный способ
func intersec1(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for elem1 := range set1 {
		for elem2 := range set2 {
			if elem1 == elem2 {
				result[elem1] = struct{}{}
			}
		}
	}

	return result
}

func intersec2(set1, set2 map[int]struct{}) map[int]struct{} {
	result := make(map[int]struct{})

	for elem := range set1 {
		// проверка, существует ли element во втором множестве set2
		_, ok := set2[elem]
		// If the key exists
		if ok {
			result[elem] = struct{}{}
		}
	}

	return result
}

func main() {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	set1[1] = struct{}{}
	set1[2] = struct{}{}
	set1[3] = struct{}{}

	set2[2] = struct{}{}
	set2[3] = struct{}{}
	set2[4] = struct{}{}

	intersec := intersec2(set1, set2)

	fmt.Println(intersec)
}
