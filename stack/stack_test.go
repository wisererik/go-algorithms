package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var s = new(Stack)

	s.Push("apple")
	s.Push("pear")
	s.Push("peach")
	s.Pop()
	s.Push("grape")
	s.Push("banana")

	assert.Equal(t, "banana", s.Top().(string))
	assert.Equal(t, 4, s.Size())
	for i := 0; i < 4; i++ {
		assert.NotEqual(t, "peach", s.Pop().(string))
	}
}
