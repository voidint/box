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

package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBPage(t *testing.T) {
	for _, item := range []struct {
		InPageNo       uint64
		InPageSize     uint64
		InTotalRecords uint64
		OutLimit       uint64
		OutOffset      uint64
		OutTotalPages  uint64
	}{
		{InPageNo: 1, InPageSize: 10, InTotalRecords: 11, OutLimit: 10, OutOffset: 0, OutTotalPages: 2},
		{InPageNo: 2, InPageSize: 3, InTotalRecords: 10, OutLimit: 3, OutOffset: 3, OutTotalPages: 4},
		{InPageNo: 0, InPageSize: 0, InTotalRecords: 10, OutLimit: 10, OutOffset: 0, OutTotalPages: 1},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 100, OutLimit: 20, OutOffset: 0, OutTotalPages: 5},
		{InPageNo: 1, InPageSize: 30, InTotalRecords: 22, OutLimit: 20, OutOffset: 0, OutTotalPages: 2},
	} {
		pg := NewPage[uint64](item.InPageNo, item.InPageSize, WithMaxPageSize(uint64(20)))
		assert.Equal(t, item.OutOffset, pg.Offset())
		assert.Equal(t, item.OutLimit, pg.Limit())
		assert.Equal(t, item.OutTotalPages, pg.TotalPages(item.InTotalRecords))
	}
}
