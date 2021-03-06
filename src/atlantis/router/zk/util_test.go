/* Copyright 2014 Ooyala, Inc. All rights reserved.
 *
 * This file is licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is
 * distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and limitations under the License.
 */

package zk

import (
	"testing"
)

func TestArrayDiff(t *testing.T) {
	arr0 := []string{}
	arr1 := []string{"mary", "had", "a"}
	arr2 := []string{"mary", "had", "a", "little", "lamb"}

	if len(ArrayDiff(arr0, arr1)) != 0 {
		t.Errorf("should be empty")
	}

	if len(ArrayDiff(arr1, arr0)) != 3 {
		t.Errorf("should be {mary, had, a}")
	}

	if len(ArrayDiff(arr2, arr1)) != 2 {
		t.Errorf("should be {little, lamb}")
	}
}
