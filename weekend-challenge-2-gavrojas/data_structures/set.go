package datastructures

import "fmt"

/* A set is a data structure that stores unique elements
 * of the same type in a sorted order. Each value is a key,
 * which means that we access each value using the value itself.
 * Accordingly, each value in a set must be unique.
 */
type Set struct {
	// A common way to configure a set is to have a key with desired store value and a boolean indicating it exists
	Items map[string]bool
}

func (s *Set) Append(value string) error {
	// What's the current state of items?
	if s.Exists(value) {
		return fmt.Errorf("item already exists")
	} // Initialize if necessary
	if s.Items == nil {
		s.Items = make(map[string]bool)
	}
	s.Items[value] = true
	return nil
}

func (s Set) Exists(value string) bool {
	// We can get a key from the map, remember the return come as (value, exists)
	_, exists := s.Items[value]
	return exists
}

func (s *Set) Delete(value string) {
	// Remember there is a method called delete, how can you use it here?
	delete(s.Items, value)
}
