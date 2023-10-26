package multiset

import (
	"fmt"
	"sort"
	"strings"
)

type Multiset[T comparable] struct {
	items map[T]int
	count int
}

// New contructs an empty multiset.
//
// The multiset is created with an initial capacity of 0.
// The initial capacity does not bound its size,
// and it grows to accomodate the number of values stored in it.
func New[T comparable]() *Multiset[T] {
	return &Multiset[T]{items: make(map[T]int), count: 0}
}

// WithCapacity constructs an empty multiset with the specified capacity.
//
// The multiset is created with an initial capacity of n.
// The initial capacity does not bound its size,
// and it grows to accomodate the number of values stored in it.
func WithCapacity[T comparable](n int) *Multiset[T] {
	return &Multiset[T]{items: make(map[T]int, n), count: 0}
}

// Insert inserts a new value v to multiset m.
//
// Insert returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Insert(v T) int {
	return m.InsertMany(v, 1)
}

// InsertMany inserts value v to multiset m, with n number of occurences.
//
// InsertMany returns the number of occurences of value v previously
// in multiset m.
func (m *Multiset[T]) InsertMany(v T, n int) int {
	if n == 0 {
		return m.Contains(v)
	}

	m.count += n
	pn := m.items[v]
	m.items[v] = pn + n

	return pn
}

// Union constructs a new multiset union of multiset m and other.
//
// The resulting multiset is a multiset of the maximum multiplicity
// of items present in m and other.
func (m *Multiset[T]) Union(other *Multiset[T]) *Multiset[T] {
	result := m.Clone()
	if other == nil || other.IsEmpty() {
		return result
	}

	other.Each(func(otherV T, otherN int) bool {
		if n, ok := result.items[otherV]; !ok || otherN > n {
			result.count -= n
			result.items[otherV] = otherN
			result.count += otherN
		}
		return false
	})

	return result
}

// Intersection constructs a new multiset intersection of multiset m and other.
//
// The resulting multiset is a multiset of the minimum multiplicity
// of items present in m and other.
func (m *Multiset[T]) Intersection(other *Multiset[T]) *Multiset[T] {
	if other == nil || other.IsEmpty() {
		return New[T]()
	}

	result := WithCapacity[T](max(len(m.items), len(other.items)))
	m.Each(func(v T, n int) bool {
		if otherN, ok := other.items[v]; ok {
			newN := min(n, otherN)
			result.items[v] = newN
			result.count += newN
		}
		return false
	})

	return result
}

// Sum constructs a new multiset sum of multiset m and other.
//
// The resulting multiset is a multiset whose multiplicities represents
// how many times a given item occur in both m and other.
func (m *Multiset[T]) Sum(other *Multiset[T]) *Multiset[T] {
	result := m.Clone()
	if other == nil || other.IsEmpty() {
		return result
	}

	other.Each(func(otherV T, otherN int) bool {
		result.InsertMany(otherV, otherN)
		return false
	})

	return result
}

// Difference constructs a new multiset difference of multiset m and other.
//
// The resulting multiset is a multiset whose multiplicities represents
// how many more of a given item are in m than other.
func (m *Multiset[T]) Difference(other *Multiset[T]) *Multiset[T] {
	if other == nil || other.IsEmpty() {
		return m.Clone()
	}

	result := WithCapacity[T](len(m.items))
	m.Each(func(v T, n int) bool {
		otherN, ok := other.items[v]
		newN := n - otherN
		if ok && newN > 0 || !ok {
			result.items[v] = newN
			result.count += newN
		}
		return false
	})

	return result
}

// Replace replaces all existing occurences of value v in multiset m, if any, with 1.
//
// Replace returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Replace(v T) int {
	n := m.items[v]
	m.count -= n
	m.items[v] = 1
	m.count += 1
	return n
}

// Remove removes value v from multiset m.
//
// Remove returns the number of occurences of value v previously in
// multiset m.
func (m *Multiset[T]) Remove(v T) int {
	n, ok := m.items[v]
	if !ok {
		return 0
	}

	m.count -= 1
	if n == 1 {
		delete(m.items, v)
	} else {
		m.items[v] = n - 1
	}

	return n
}

// Get returns a value from multiset m that equals value v,
// along with its number of occurences and a boolean that is true.
//
// Get returns the zero value of type T, count 0 and false if value v
// does not exist in multiset m.
func (m *Multiset[T]) Get(v T) (T, int, bool) {
	for val, c := range m.items {
		if v == val {
			return val, c, true
		}
	}
	return *new(T), 0, false
}

// Contains returns the number of occurences of value v in multiset m.
func (m *Multiset[T]) Contains(v T) int {
	if n, ok := m.items[v]; ok {
		return n
	}
	return 0
}

// IsEmpty returns true if there are no items in multiset m,
// otherwise false.
func (m *Multiset[T]) IsEmpty() bool {
	return m.count == 0
}

// Len returns the number of items of multiset m.
//
// Duplicates are counted.
func (m *Multiset[T]) Len() int {
	return m.count
}

// Cardinality returns the number of items in multiset m.
//
// Multiplicity of an item is not considered.
func (m *Multiset[T]) Cardinality() int {
	return len(m.items)
}

// Each iterates over all items and calls f for each item present in
// multiset m.
//
// If f returns true, Each stops the iteration.
func (m *Multiset[T]) Each(f func(T, int) bool) {
	for v, n := range m.items {
		if f(v, n) {
			break
		}
	}
}

// Clone returns a new multiset which is a copy of multiset m.
func (m *Multiset[T]) Clone() *Multiset[T] {
	result := WithCapacity[T](m.Cardinality())
	m.Each(func(v T, n int) bool {
		result.InsertMany(v, n)
		return false
	})
	return result
}

// Equal returns true if the length and items of
// multiset m and other are equal.
//
// TODO: Update this when https://github.com/golang/go/issues/57436 lands.
func (m *Multiset[T]) Equal(other *Multiset[T]) bool {
	if m.Len() != other.Len() || len(m.items) != len(other.items) {
		return false
	}

	for v, n := range m.items {
		if otherN, ok := other.items[v]; !ok || n != otherN {
			return false
		}
	}

	return true
}

// String returns a formatted multiset with the following format:
//
// Multiset{1: 2, 2: 3}
func (m *Multiset[T]) String() string {
	items := make([]string, 0, len(m.items))

	m.Each(func(v T, n int) bool {
		items = append(items, fmt.Sprintf("%v:%d", v, n))
		return false
	})

	sort.Strings(items)
	return fmt.Sprintf("Multiset{%s}", strings.Join(items, ", "))
}

// max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// min returns the smaller of x or y.
func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m *Multiset[T]) Print() {
	for item, count := range m.items {
		for i := 0; i < count; i++ {
			fmt.Println(item)
		}
	}
}