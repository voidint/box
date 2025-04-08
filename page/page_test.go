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

package page

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPage(t *testing.T) {
	for _, item := range []struct {
		PageNo       uint64
		PageSize     uint64
		TotalRecords uint64
		TotalPages   uint64
		Offset       uint64
		Records      []uint64
	}{
		{PageNo: 1, PageSize: 10, TotalRecords: 100, TotalPages: 10, Offset: 0},
		{PageNo: 2, PageSize: 3, TotalRecords: 7, TotalPages: 3, Offset: 3},
	} {
		pgr := NewPager[uint64](item.PageNo, item.PageSize, item.TotalRecords)

		pg := pgr.BuildPage()
		assert.NotNil(t, pg)
		assert.Equal(t, item.TotalPages, pg.TotalPages)

		dbPg := pgr.BuildDBPage()
		assert.NotNil(t, dbPg)
		assert.Equal(t, item.PageSize, dbPg.Limit())
		assert.Equal(t, item.Offset, dbPg.Offset())
		assert.Equal(t, item.PageNo, dbPg.PageNo())
		assert.Equal(t, item.PageSize, dbPg.PageSize())
	}
}

func TestMustCalculateTotalPages(t *testing.T) {
	for _, item := range []struct {
		PageSize     uint64
		TotalRecords uint64
		TotalPages   uint64
		Panic        bool
	}{
		{PageSize: 10, TotalRecords: 10, TotalPages: 1},
		{PageSize: 10, TotalRecords: 0, TotalPages: 0},
		{PageSize: 3, TotalRecords: 10, TotalPages: 4},
		{PageSize: 3, TotalRecords: 0, TotalPages: 0},
		{PageSize: 0, TotalRecords: 10, TotalPages: 0, Panic: true},
	} {
		if item.Panic {
			assert.Panics(t, func() {
				mustCalculateTotalPages(item.PageSize, item.TotalRecords)
			})
		} else {
			mustCalculateTotalPages(item.PageSize, item.TotalRecords)
		}
	}
}
