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
