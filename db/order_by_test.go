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

func TestOneOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field     string
		Direction OrderByDirection
		Out1      string
		Out2      []string
	}{
		{
			Field:     "age",
			Direction: DESC,
			Out1:      "age DESC",
			Out2:      []string{"age DESC"},
		},
		{
			Field:     "`age`",
			Direction: ASC,
			Out1:      "`age` ASC",
			Out2:      []string{"`age` ASC"},
		},
		{
			Field:     "",
			Direction: ASC,
			Out1:      " ASC",
			Out2:      []string{" ASC"},
		},
	} {
		assert.Equal(t, item.Out1, OneOrderBy(item.Field, item.Direction).String())
		assert.Equal(t, item.Out2, OneOrderBy(item.Field, item.Direction).Strings())
	}
}

func TestTwoOrderBy(t *testing.T) {
	for _, item := range []struct {
		Field1     string
		Direction1 OrderByDirection
		Field2     string
		Direction2 OrderByDirection
		Out1       string
		Out2       []string
	}{
		{
			Field1:     "age",
			Direction1: DESC,
			Field2:     "gender",
			Direction2: ASC,
			Out1:       "age DESC,gender ASC",
			Out2:       []string{"age DESC", "gender ASC"},
		},
		{
			Field1:     "`age`",
			Direction1: DESC,
			Field2:     "gender",
			Direction2: ASC,
			Out1:       "`age` DESC,gender ASC",
			Out2:       []string{"`age` DESC", "gender ASC"},
		},
		{
			Field1:     "`age`",
			Direction1: ASC,
			Out1:       "`age` ASC, ",
			Out2:       []string{"`age` ASC", " "},
		},
	} {
		assert.Equal(t, item.Out1, TwoOrderBy(item.Field1, item.Direction1, item.Field2, item.Direction2).String())
		assert.Equal(t, item.Out2, TwoOrderBy(item.Field1, item.Direction1, item.Field2, item.Direction2).Strings())
	}
}
