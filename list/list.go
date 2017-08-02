package list

import (
	"fmt"
)

type List interface {
	Add(values ...interface{})

	Insert(index int, values ...interface{})

	Get(index int) (interface{}, bool)

	Remove(index int) bool

	Contains(value interface{}) bool

	Empty() bool

	Size() int

	Values() []interface{}
}

const (
	extendFactor = float32(2.0)
)

type ArrayList struct {
	elements []interface{}
	size     int
}

func (l *ArrayList) Empty() bool {
	return l.size == 0
}

func (l *ArrayList) Size() int {
	return l.size
}

func (l *ArrayList) Add(values ...interface{}) {
	l.extendCapacity(len(values))
	for _, value := range values {
		l.elements[l.size] = value
		l.size++
	}

}

func (l *ArrayList) Insert(index int, values ...interface{}) {
	if index > l.size {
		return
	}

	if index == l.size {
		l.Add(values)
	}

	length := len(values)
	l.size += length
	l.extendCapacity(length)

	fmt.Println(l.size)
	for i := l.size - 1; i >= index+length; i-- {

		l.elements[i] = l.elements[i-length]
	}

	for i, value := range values {
		l.elements[index+i] = value
	}
}

func (l *ArrayList) Get(index int) (interface{}, bool) {
	if index >= 0 && index < l.size {
		return l.elements[index], true
	}

	return nil, false
}

func (l *ArrayList) Contains(value interface{}) bool {
	flag := false
	for _, searchedValue := range l.elements {
		if value == searchedValue {
			flag = true
			break
		}
	}
	return flag
}

func (l *ArrayList) Remove(index int) bool {
	if index >= 0 && index < l.size {
		l.elements[index] = nil
		copy(l.elements[index:], l.elements[index+1:l.size])
		l.size--
		return true
	}

	return false
}

func (l *ArrayList) Values() []interface{} {
	return l.elements[:l.size]
}

func (l *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, l.elements)
	l.elements = newElements
}

func (l *ArrayList) extendCapacity(n int) {

	if current := cap(l.elements); current <= l.size+n {
		newCap := int(extendFactor * float32(current+n))
		l.resize(newCap)
	}
}
