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

func TestInSlice(t *testing.T) {
	assert.True(t, InSlice("voidint", []string{"hello", "world", "voidint"}))
	assert.False(t, InSlice("voidint", []string{"hello", "world"}))
	assert.True(t, InSlice(7, []int{1, 2, 3, 4, 5, 6, 7}))
	assert.False(t, InSlice(77, []int{1, 2, 3, 4, 5, 6, 7}))
}
