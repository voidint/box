package condexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAny(t *testing.T) {
	t.Run("整型", func(t *testing.T) {
		assert.Equal(t, 2, Any(1 > 2, 1, 2))
		assert.Equal(t, 2, Any(1 > 2, 1, 2))
	})

	t.Run("浮点型", func(t *testing.T) {
		assert.Equal(t, float32(2.0), Any[float32](1.0 > 2.0, 1.0, 2.0))
		assert.Equal(t, float64(2.0), Any(1.0 > 2.0, 1.0, 2.0))
	})

	t.Run("字符串型", func(t *testing.T) {
		assert.Equal(t, "1<=2", Any(1 > 2, "1>2", "1<=2"))
		assert.Equal(t, "1<=2", Any(1 > 2, "1>2", "1<=2"))
	})
}
