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

var setValue = struct{}{}

// Set 由不重复的T组成的集合。该集合非线程安全。
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet 返回指定容量的Set集合。
func NewSet[T comparable](capacity int, elements ...T) *Set[T] {
	var items map[T]struct{}
	if capacity <= 0 {
		capacity = len(elements)
	}
	items = make(map[T]struct{}, capacity)
	for i := range elements {
		items[elements[i]] = setValue
	}
	return &Set[T]{
		items: items,
	}
}

// Add 往Set集合中添加元素。
func (set *Set[T]) Add(elements ...T) *Set[T] {
	for i := range elements {
		set.items[elements[i]] = setValue
	}
	return set
}

// Remove 移除元素
func (set *Set[T]) Remove(elements ...T) *Set[T] {
	for i := range elements {
		delete(set.items, elements[i])
	}
	return set
}

// Elements 返回当前Set集合中的所有元素。
func (set *Set[T]) Elements() []T {
	all := make([]T, 0, len(set.items))
	for k := range set.items {
		all = append(all, k)
	}
	return all
}

// Contains 返回是否包含当前元素的布尔值
func (set *Set[T]) Contains(element T) bool {
	_, ok := set.items[element]
	return ok
}

// Length 返回当前的元素个数
func (set *Set[T]) Length() int {
	return len(set.items)
}

// IsEmpty 返回当前集合是否为空
func (set *Set[T]) IsEmpty() bool {
	return len(set.items) <= 0
}
