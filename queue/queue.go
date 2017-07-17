package queue

type Node struct {
	data interface{}
	next *Node
}

// FIFO
type Queue struct {
	first *Node
	last  *Node
	size  int
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) Enqueue(element interface{}) {
	node := &Node{element, nil}
	if q.last != nil {
		last := q.last
		last.next, q.last = node, node

	} else {
		q.first, q.last = node, node
	}

	q.size++
}

func (q *Queue) Dequeue() interface{} {
	if q.first != nil {
		element := q.first.data
		q.first = q.first.next
		q.size--
		return element
	}
	return nil
}

func (q *Queue) First() interface{} {
	if q.first != nil {
		return q.first.data
	}
	return nil
}
