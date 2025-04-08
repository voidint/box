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
