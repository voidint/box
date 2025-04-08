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
	"github.com/voidint/box/constraints"
)

// WithMaxPageSize 设置分页大小上限阈值
func WithMaxPageSize[INT constraints.Unsigned](maxPageSize INT) func(*Page[INT]) {
	return func(pg *Page[INT]) {
		pg.maxPageSize = maxPageSize
	}
}

// Page 分页
type Page[INT constraints.Unsigned] struct {
	pageNo, pageSize, maxPageSize INT
}

// NewPage 返回分页实例
func NewPage[INT constraints.Unsigned](pageNo, pageSize INT, opts ...func(*Page[INT])) (pg *Page[INT]) {
	pg = &Page[INT]{
		pageNo:   pageNo,
		pageSize: pageSize,
	}

	if pageNo <= 0 {
		pg.pageNo = 1
	}

	if pageSize <= 0 {
		pg.pageSize = 10
	}

	for _, setter := range opts {
		setter(pg)
	}

	if pg.maxPageSize > 0 && pg.pageSize > pg.maxPageSize {
		pg.pageSize = pg.maxPageSize
	}

	return
}

// Limit 返回分页大小
func (pg *Page[INT]) Limit() INT {
	return pg.pageSize
}

// Offset 返回分页记录偏移量
func (pg *Page[INT]) Offset() INT {
	return (pg.pageNo - 1) * pg.pageSize
}

// PageNo 返回分页页号
func (pg *Page[INT]) PageNo() INT {
	return pg.pageNo
}

// PageSize 返回分页大小
func (pg *Page[INT]) PageSize() INT {
	return pg.pageSize
}

// TotalPages 返回总记录数对应的总分页数
func (pg *Page[INT]) TotalPages(totalRecords INT) (totalPages INT) {
	if totalRecords <= 0 {
		return 0
	}
	if totalRecords%pg.pageSize == 0 {
		return totalRecords / pg.pageSize
	}
	return totalRecords/pg.pageSize + 1
}
