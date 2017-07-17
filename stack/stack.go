package stack

type Node struct {
	data interface{}
	next *Node
}

// FILO
type Stack struct {
	last *Node
	size int
}

func (s *Stack) Push(element interface{}) {
	s.last = &Node{element, s.last}
	s.size++
}

func (s *Stack) Pop() interface{} {
	if s.size > 0 {
		element := s.last.data
		s.last = s.last.next
		s.size--
		return element
	}
	return nil
}

func (s *Stack) Top() interface{} {
	if s.size > 0 {
		return s.last.data
	}
	return nil
}

func (s *Stack) Size() int {
	return s.size
}
