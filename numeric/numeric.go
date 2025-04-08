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
	"fmt"
	"strings"
)

// PositiveToPtr 若入参v大于零，则返回其指针；否则返回nil。
func PositiveToPtr[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64](v T) *T {
	if v > 0 {
		return &v
	}
	return nil
}

// PositiveUint8ToUint32Ptr 若参数v大于零，则将其转换成uint32后返回其指针；若参数v等于零，则返回nil。
func PositiveUint8ToUint32Ptr(v uint8) *uint32 {
	if v == 0 {
		return nil
	}
	ret := uint32(v)
	return &ret
}

// ExtractFromPtr 若入参指针为nil，则返回0；否则，返回指针指向的内存所存储的值。
func ExtractFromPtr[T ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~float32 | ~float64](p *T) T {
	if p == nil {
		return 0
	}
	return *p
}

type Number interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~int | ~float32 | ~float64
}

// Join 返回由指定分隔符所连接而成的字符串
func Join[T Number](items []T, sep string) string {
	switch len(items) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%v", items[0])
	case 2:
		return fmt.Sprintf("%v%s%v", items[0], sep, items[1])
	case 3:
		return fmt.Sprintf("%v%s%v%s%v", items[0], sep, items[1], sep, items[2])
	default:
		var buf strings.Builder
		for i := range items {
			buf.WriteString(fmt.Sprintf("%v", items[i]))
			if i != len(items)-1 {
				buf.WriteString(sep)
			}
		}
		return buf.String()
	}
}
