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

package str

// NonEmptyToPtr 若入参v非空字符串，则返回其指针；否则返回nil。
func NonEmptyToPtr(v string) *string {
	if v != "" {
		return &v
	}
	return nil
}

// ExtractFromPtr 若入参指针为nil，则返回空字符串；否则，返回指针指向的内存所存储的字符串。
func ExtractFromPtr(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
