package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	t.Run("string set", func(t *testing.T) {
		set := NewSet(3, "hello", "world")
		assert.Equal(t, 2, set.Length())

		set.Add("voidint")
		assert.Equal(t, 3, set.Length())
		assert.Contains(t, set.Elements(), "hello", "world", "voidint")
		set.Remove("voidint")
		assert.Equal(t, false, set.Contains("voidint"))
		set.Remove("hello")
		assert.Equal(t, false, set.IsEmpty())
		set.Remove("world")
		assert.Equal(t, true, set.IsEmpty())
	})
}
