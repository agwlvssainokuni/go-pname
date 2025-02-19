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
	"bytes"
	"strings"
)

type Token struct {
	Lnm string
	Pnm []string
	OK  bool
}

type Tokenizer interface {
	SplitText(text string) []*Token
}

type DictTokenizer struct {
	dict map[string][]string
}

func NewDictTokenizer(dict map[string][]string) Tokenizer {
	return &DictTokenizer{dict}
}

func (tk *DictTokenizer) SplitText(text string) []*Token {
	result := make([]*Token, 0, len(text))
	unmatch := bytes.NewBuffer(make([]byte, 0, len(text)))
	for i := 0; i < len(text); i++ {
		token := findLongestToken(text[i:], tk.dict)
		if token == nil {
			unmatch.WriteByte(text[i])
			continue
		}
		if unmatch.Len() > 0 {
			result = append(result, &Token{unmatch.String(), []string{unmatch.String()}, false})
			unmatch = bytes.NewBuffer(make([]byte, 0, len(text)))
		}
		result = append(result, token)
		i += len(token.Lnm) - 1
	}
	if unmatch.Len() > 0 {
		result = append(result, &Token{unmatch.String(), []string{unmatch.String()}, false})
	}
	return result
}

func findLongestToken(text string, dict map[string][]string) *Token {
	var (
		ln string
		pn []string
	)
	length := 0
	for l, p := range dict {
		if len(l) <= length {
			continue
		}
		if strings.HasPrefix(text, l) {
			ln = l
			pn = p
			length = len(ln)
		}
	}
	if length <= 0 {
		return nil
	}
	return &Token{ln, pn, true}
}
