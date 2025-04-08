// Copyright (c) 2025 voidint <voidint@126.com>. All rights reserved.
//
// This source code is licensed under the license found in the
// LICENSE file in the root directory of this source tree.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
