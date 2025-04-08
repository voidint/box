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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractFromPtr(t *testing.T) {
	s1 := ""
	s2 := "box"

	type args struct {
		p *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "入参是nil",
			args: args{p: nil},
			want: "",
		},
		{
			name: "入参是空字符串",
			args: args{p: &s1},
			want: s1,
		},
		{
			name: "入参是非空字符串",
			args: args{p: &s2},
			want: s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, ExtractFromPtr(tt.args.p))
		})
	}
}

func TestNonEmptyToPtr(t *testing.T) {
	s1 := " 	"
	s2 := "box"

	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "入参是空字符串",
			args: args{v: ""},
			want: nil,
		},
		{
			name: "入参是空白字符串",
			args: args{v: s1},
			want: &s1,
		},
		{
			name: "入参是非空字符串",
			args: args{v: s2},
			want: &s2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NonEmptyToPtr(tt.args.v))
		})
	}
}
