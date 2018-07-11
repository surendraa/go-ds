package list

type single interface {
	New([]interface{}) *Single
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
func (s *Single) New(items []interface{}) *Single {
	if len(items) > 0 {
		s.Head = &node{items[0], nil}
		cur := s.Head
		for i := range items[1:] {
			cur.Next = &node{i, nil}
		}
	} else {
		s.Head = &node{}
	}
	return s
}
