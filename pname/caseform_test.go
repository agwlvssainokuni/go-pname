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

import "testing"

func TestToLowerCamelCase(t *testing.T) {
	assertString(t, "", ToLowerCamelCase([]string{""}))
	assertString(t, "abc", ToLowerCamelCase([]string{"ABC"}))
	assertString(t, "abcDefGhi", ToLowerCamelCase([]string{"ABC", "DEF", "GHI"}))
}

func TestToUpperCamelCase(t *testing.T) {
	assertString(t, "", ToUpperCamelCase([]string{""}))
	assertString(t, "Abc", ToUpperCamelCase([]string{"ABC"}))
	assertString(t, "AbcDefGhi", ToUpperCamelCase([]string{"ABC", "DEF", "GHI"}))
}

func TestToLowerSnakeCase(t *testing.T) {
	assertString(t, "", ToLowerSnakeCase([]string{""}))
	assertString(t, "abc", ToLowerSnakeCase([]string{"ABC"}))
	assertString(t, "abc_def_ghi", ToLowerSnakeCase([]string{"ABC", "DEF", "GHI"}))
}

func TestToUpperSnakeCase(t *testing.T) {
	assertString(t, "", ToUpperSnakeCase([]string{""}))
	assertString(t, "ABC", ToUpperSnakeCase([]string{"abc"}))
	assertString(t, "ABC_DEF_GHI", ToUpperSnakeCase([]string{"abc", "def", "ghi"}))
}
