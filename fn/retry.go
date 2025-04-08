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

package fn

import "time"

// Retry Retry function call
func Retry(retries int, retryDelay time.Duration, do func() (again bool, err error)) (err error) {
	var again bool

	if retries < 0 {
		for {
			if again, err = do(); !again {
				break
			}
			time.Sleep(retryDelay)
		}

	} else {
		for i := 0; i <= retries; i++ {
			if again, err = do(); !again {
				break
			}
			time.Sleep(retryDelay)
		}
	}

	return err
}
