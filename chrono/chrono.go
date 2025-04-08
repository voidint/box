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

package chrono

import (
	"time"

	"github.com/dromara/carbon/v2"
)

// LastFewDays 返回最近N天的时间列表
func LastFewDays(n int) (days []carbon.Carbon) {
	if n <= 0 {
		panic("invalid parameters")
	}

	days = make([]carbon.Carbon, 0, n)

	today := carbon.Now().StartOfDay()
	for i := n - 1; i > 0; i-- {
		days = append(days, today.SubDays(i))
	}
	days = append(days, today)

	return days
}

// Last7Days 返回最近7天的时间列表
func Last7Days() (days []carbon.Carbon) {
	return LastFewDays(7)
}

// Last30Days 返回最近30天的时间列表
func Last30Days() (days []carbon.Carbon) {
	return LastFewDays(30)
}

// WithinMonth 返回指定年月的起止时间
func WithinMonth(year int, month int) (start, end carbon.Carbon) {
	if month < 1 || month > 12 {
		panic("invalid month")
	}
	c := carbon.CreateFromDate(year, month, 1)
	return c.StartOfMonth(), c.EndOfMonth()
}

// WithinAFewDays 返回若干天内的起止时间
func WithinAFewDays(n int) (start, end carbon.Carbon) {
	now := carbon.Now()
	return now.SubDays(n).StartOfDay(), now.EndOfDay()
}

// Within7Days 返回'7天内'的开始时间和截止时间
func Within7Days() (start, end carbon.Carbon) {
	return WithinAFewDays(7)
}

// Within30Days 返回'30天内'的开始时间和截止时间
func Within30Days() (start, end carbon.Carbon) {
	return WithinAFewDays(30)
}

// IsItToday 返回是否是今天的布尔值
func IsItToday(t time.Time) bool {
	now := time.Now()
	return t.Day() == now.Day() && t.Month() == now.Month() && t.Year() == now.Year()
}
