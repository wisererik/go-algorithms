package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	var list List = new(ArrayList)

	assert.Equal(t, 0, list.Size())
	assert.Equal(t, true, list.Empty())

	list.Add("apple", "pear", "grape", "banana")

	assert.Equal(t, 4, list.Size())
	assert.Equal(t, false, list.Empty())
	assert.Equal(t, true, list.Contains("apple"))
	assert.Equal(t, false, list.Contains("empty"))

	list.Insert(8, "monkey")
	assert.Equal(t, 4, list.Size())

	list.Insert(1, "monkey", "panda")
	expected := []interface{}{"apple", "monkey", "panda", "pear", "grape", "banana"}

	assert.Equal(t, expected, list.Values())

}
