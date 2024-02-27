package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	s := NewQueue[int]()

	assert := assert.New(t)
	assert.True(s.IsEmpty())

	s.EnqueueAll(1, 2, 3, 4, 5, 6)

	assert.False(s.IsEmpty())

	for i := 1; i < 7; i++ {
		val, err := s.Dequeue()
		assert.Nil(err)
		assert.Equal(val, i)
	}

	_, err := s.Dequeue()
	assert.Error(err)
}
