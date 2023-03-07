package numeric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFromPtr(t *testing.T) {
	v1 := uint8(1)
	v2 := uint16(2)
	v3 := uint32(3)
	v4 := uint64(4)
	v5 := uint(5)
	v6 := float32(6)
	v7 := float64(7)

	assert.Equal(t, uint8(0), ExtractFromPtr[uint8](nil))
	assert.Equal(t, v1, ExtractFromPtr(&v1))
	assert.Equal(t, v2, ExtractFromPtr(&v2))
	assert.Equal(t, v3, ExtractFromPtr(&v3))
	assert.Equal(t, v4, ExtractFromPtr(&v4))
	assert.Equal(t, v5, ExtractFromPtr(&v5))
	assert.Equal(t, v6, ExtractFromPtr(&v6))
	assert.Equal(t, v7, ExtractFromPtr(&v7))
}
