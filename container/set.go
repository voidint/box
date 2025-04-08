// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package container

var setValue = struct{}{}

// Set implements a generic collection of unique elements with type safety.
// The zero value is not usable, always create instances via NewSet constructor.
// This implementation is not thread-safe and requires external synchronization
// when used in concurrent environments.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet constructs a Set instance with optional initial capacity and elements.
func NewSet[T comparable](capacity int, elements ...T) *Set[T] {
	var items map[T]struct{}
	if capacity <= 0 {
		capacity = len(elements)
	}
	items = make(map[T]struct{}, capacity)
	for i := range elements {
		items[elements[i]] = setValue
	}
	return &Set[T]{
		items: items,
	}
}

// Add inserts one or multiple elements into the set.
func (set *Set[T]) Add(elements ...T) *Set[T] {
	for i := range elements {
		set.items[elements[i]] = setValue
	}
	return set
}

// Remove deletes specified elements from the set.
func (set *Set[T]) Remove(elements ...T) *Set[T] {
	for i := range elements {
		delete(set.items, elements[i])
	}
	return set
}

// Elements returns all set members as a new slice.
func (set *Set[T]) Elements() []T {
	all := make([]T, 0, len(set.items))
	for k := range set.items {
		all = append(all, k)
	}
	return all
}

// Contains checks membership of an element in the set.
func (set *Set[T]) Contains(element T) bool {
	_, ok := set.items[element]
	return ok
}

// Length returns the cardinality (element count) of the set.
func (set *Set[T]) Length() int {
	return len(set.items)
}

// IsEmpty checks for an empty set condition.
func (set *Set[T]) IsEmpty() bool {
	return len(set.items) <= 0
}
