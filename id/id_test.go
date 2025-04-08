// Copyright (c) 2025 voidint <voidint@126.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultNode(t *testing.T) {
	t.Run("Default node generates IDs", func(t *testing.T) {
		t.Run("String type ID", func(t *testing.T) {
			assert.Equal(t, 19, len(String()))
			for i := 0; i < 1000; i++ {
				assert.NotEqual(t, String(), String())
			}
		})
		t.Run("Int64 type ID", func(t *testing.T) {
			assert.Equal(t, true, Int64() > 0)
			for i := 0; i < 100; i++ {
				assert.NotEqual(t, Int64(), Int64())
			}
		})
	})
}

func TestNewNode(t *testing.T) {
	t.Run("Custom node generates IDs", func(t *testing.T) {
		t.Run("Node instantiation", func(t *testing.T) {
			node1, err := NewNode(1)
			assert.Nil(t, err)
			assert.NotNil(t, node1)

			node1023, err := NewNode(1023)
			assert.Nil(t, err)
			assert.NotNil(t, node1023)

			node1024, err := NewNode(1024)
			assert.NotNil(t, err)
			assert.Nil(t, node1024)

			assert.NotNil(t, MustNewNode(1))
			assert.Panics(t, func() { MustNewNode(1024) })
		})

		t.Run("ID generation", func(t *testing.T) {
			node := MustNewNode(1023)
			assert.NotEmpty(t, node.String())
			for i := 0; i < 1000; i++ {
				assert.NotEqual(t, node.String(), node.String())
			}
			assert.Greater(t, node.Int64(), int64(0))
			for i := 0; i < 1000; i++ {
				assert.NotEqual(t, node.Int64(), node.Int64())
			}
		})
	})
}
