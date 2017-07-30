/*
 * Copyright 2017 agwlvssainokuni
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package pname

import (
	"testing"
)

func assertString(t *testing.T, expected, result string) {
	if expected != result {
		t.Errorf("should be %s, but %s", expected, result)
	}
}

func assertInt(t *testing.T, expected, result int) {
	if expected != result {
		t.Errorf("should be %v, but %v", expected, result)
	}
}

func assertBool(t *testing.T, expected, result bool) {
	if expected != result {
		t.Errorf("should be %v, but %v", expected, result)
	}
}
