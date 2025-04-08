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

package id

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultNode(t *testing.T) {
	t.Run("默认ID节点生成ID", func(t *testing.T) {
		t.Run("string类型ID", func(t *testing.T) {
			assert.Equal(t, 19, len(String()))
			for i := 0; i < 1000; i++ {
				assert.NotEqual(t, String(), String())
			}
		})
		t.Run("int64类型ID", func(t *testing.T) {
			assert.Equal(t, true, Int64() > 0)
			for i := 0; i < 100; i++ {
				assert.NotEqual(t, Int64(), Int64())
			}
		})
	})
}

func TestNewNode(t *testing.T) {
	t.Run("自定义节点生成ID", func(t *testing.T) {
		t.Run("节点实例化", func(t *testing.T) {
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

		t.Run("生成ID", func(t *testing.T) {
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
