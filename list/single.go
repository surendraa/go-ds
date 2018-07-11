package list

import (
	"errors"
	"reflect"
)

type single interface {
	New([]interface{}) (*Single, error)
	// InsertFront()
	// RemoveFront()
	// InsertBack()
	// RemoveBack()
	// Search(interface{})
	// IsEmpty() bool
	// InsertBefore(interface{})
	// RemoveBefore(interface{})
	// RemoveAfter(interface{})
	// InsertAfter(interface{})
	// Clear()
}

// Node : each element of the single linked list
type node struct {
	Value interface{}
	Next  *node
}

// Single : Single linked List
type Single struct {
	Head *node
}

// New : creates a single linked list
func (s *Single) New(items interface{}) (*Single, error) {
	vals := reflect.ValueOf(items)
	if vals.Kind() != reflect.Slice {
		return s, errors.New("Arg not a slice")
	}

	values := make([]interface{}, vals.Len())
	for i := 0; i < vals.Len(); i++ {
		values[i] = vals.Index(i).Interface()
	}

	if len(values) > 0 {
		s.Head = &node{values[0], nil}
		cur := s.Head
		for i := range values[1:] {
			cur.Next = &node{i, nil}
		}
	} else {
		s.Head = &node{}
	}
	return s, nil
}
