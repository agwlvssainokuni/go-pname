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
	"strings"
)

func ToLowerCamelCase(text []string) string {
	first := true
	r := make([]string, 0, len(text))
	for _, t := range text {
		if first {
			r = append(r, strings.ToLower(t))
			first = false
		} else {
			r = append(r, strings.Title(strings.ToLower(t)))
		}
	}
	return strings.Join(r, "")
}

func ToUpperCamelCase(text []string) string {
	r := make([]string, 0, len(text))
	for _, t := range text {
		r = append(r, strings.Title(strings.ToLower(t)))
	}
	return strings.Join(r, "")
}

func ToLowerSnakeCase(text []string) string {
	r := make([]string, 0, len(text))
	for _, t := range text {
		r = append(r, strings.ToLower(t))
	}
	return strings.Join(r, "_")
}

func ToUpperSnakeCase(text []string) string {
	r := make([]string, 0, len(text))
	for _, t := range text {
		r = append(r, strings.ToUpper(t))
	}
	return strings.Join(r, "_")
}
