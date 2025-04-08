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

package fn

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	var ErrRecordNotFound = errors.New("record not found")
	delay := time.Microsecond * 10

	t.Run("retries < 0", func(t *testing.T) {
		calls, retries := 0, -1
		err := Retry(retries, delay, func() (again bool, err error) {
			calls++

			if calls <= 6 {
				return true, nil
			}
			return false, ErrRecordNotFound
		})

		assert.Equal(t, ErrRecordNotFound, err)
		assert.Equal(t, 7, calls)
	})

	t.Run("retries == 0", func(t *testing.T) {
		calls, retries := 0, 0
		err := Retry(retries, delay, func() (again bool, err error) {
			calls++

			if calls <= retries {
				return true, nil
			}
			return false, ErrRecordNotFound
		})

		assert.Equal(t, ErrRecordNotFound, err)
		assert.Equal(t, retries+1, calls)
	})

	t.Run("retries > 0", func(t *testing.T) {
		calls, retries := 0, 3
		err := Retry(retries, delay, func() (again bool, err error) {
			calls++

			if calls <= retries {
				return true, nil
			}
			return false, ErrRecordNotFound
		})

		assert.Equal(t, ErrRecordNotFound, err)
		assert.Equal(t, retries+1, calls)
	})
}
