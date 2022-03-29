package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInSlice(t *testing.T) {
	assert.True(t, InSlice("voidint", []string{"hello", "world", "voidint"}))
	assert.False(t, InSlice("voidint", []string{"hello", "world"}))
	assert.True(t, InSlice(7, []int{1, 2, 3, 4, 5, 6, 7}))
	assert.False(t, InSlice(77, []int{1, 2, 3, 4, 5, 6, 7}))
}
