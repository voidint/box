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

func TestJoin(t *testing.T) {
	assert.Equal(t, "", Join([]uint64{}, "|"))
	assert.Equal(t, "1", Join([]uint64{1}, "|"))
	assert.Equal(t, "1|2", Join([]uint64{1, 2}, "|"))
	assert.Equal(t, "1|2|3", Join([]uint64{1, 2, 3}, "|"))
	assert.Equal(t, "1|2|3|4", Join([]uint64{1, 2, 3, 4}, "|"))
	assert.Equal(t, "1234", Join([]uint64{1, 2, 3, 4}, ""))
	assert.Equal(t, "1 2 3 4", Join([]uint64{1, 2, 3, 4}, " "))
	assert.Equal(t, "1|2|3|4", Join([]float64{1, 2, 3, 4}, "|"))
	assert.Equal(t, "1.12|2.23|3.34|4.45", Join([]float64{1.12, 2.23, 3.34, 4.45}, "|"))
}
