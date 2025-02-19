/*
 * Copyright 2017,2025 agwlvssainokuni
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

func createDict() map[string][]string {
	dict := make(map[string][]string)
	dict["a"] = []string{"A"}
	dict["aa"] = []string{"AA"}
	dict["b"] = []string{"B"}
	dict["bb"] = []string{"BB"}
	return dict
}

func TestFindLongestToken(t *testing.T) {
	dict := createDict()
	assertFindLongestToken(t, dict, "abb", "a")
	assertFindLongestToken(t, dict, "aab", "aa")
	assertFindLongestToken(t, dict, "aaa", "aa")
	assertFindLongestToken(t, dict, "baa", "b")
	assertFindLongestToken(t, dict, "bba", "bb")
	assertFindLongestToken(t, dict, "bbb", "bb")
	r := findLongestToken("cab", dict)
	if r != nil {
		t.Errorf("should be nil for %s", "cab")
	}
}

func TestSplitText(t *testing.T) {
	dictTk := NewDictTokenizer(createDict())
	tk1 := dictTk.SplitText("aaabbbccabc")
	assertInt(t, 8, len(tk1))
	assertString(t, "aa", tk1[0].Lnm)
	assertBool(t, true, tk1[0].OK)
	assertString(t, "a", tk1[1].Lnm)
	assertBool(t, true, tk1[1].OK)
	assertString(t, "bb", tk1[2].Lnm)
	assertBool(t, true, tk1[2].OK)
	assertString(t, "b", tk1[3].Lnm)
	assertBool(t, true, tk1[3].OK)
	assertString(t, "cc", tk1[4].Lnm)
	assertBool(t, false, tk1[4].OK)
	assertString(t, "a", tk1[5].Lnm)
	assertBool(t, true, tk1[5].OK)
	assertString(t, "b", tk1[6].Lnm)
	assertBool(t, true, tk1[6].OK)
	assertString(t, "c", tk1[7].Lnm)
	assertBool(t, false, tk1[7].OK)
}

func assertFindLongestToken(t *testing.T, dict map[string][]string, text string, expected string) {
	r := findLongestToken(text, dict)
	if r.Lnm != expected {
		t.Errorf("should be %s for %s", expected, text)
	}
}
