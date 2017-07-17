package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	var q = new(Queue)

	q.Enqueue("apple")
	q.Enqueue("pear")
	q.Enqueue("peach")
	q.Dequeue()
	q.Enqueue("grape")
	q.Enqueue("banana")

	assert.Equal(t, "pear", q.First().(string))
	assert.Equal(t, 4, q.Size())
	for i := 0; i < 4; i++ {
		assert.NotEqual(t, "apple", q.Dequeue().(string))
	}
}
