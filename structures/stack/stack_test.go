package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackCreation(t *testing.T) {
	s := NewStack[int]()

	assert := assert.New(t)
	assert.True(s.IsEmpty())
	assert.True(s.Len() == 0)
	assert.True(s.top == nil)

	_, err := s.Peek()
	assert.Error(err)

	err = s.Pop()
	assert.Error(err)
}

func TestStack(t *testing.T) {
	s := NewStack[int]()

	s.PushAll(1, 2, 3, 4, 5, 6)

	assert := assert.New(t)
	assert.False(s.IsEmpty())

	val, err := s.Peek()
	assert.Nil(err)
	assert.Equal(val, 6)

	err = s.Pop()
	assert.Nil(err)

	val, err = s.Peek()
	assert.Nil(err)
	assert.Equal(val, 5)

	s.Push(110)

	val, err = s.Peek()
	assert.Nil(err)
	assert.Equal(val, 110)
}
