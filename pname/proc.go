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
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func Process(csvr *csv.Reader, csvw *csv.Writer, tokenizer Tokenizer, getPname func([]*Token) string, getDesc func([]*Token) string) (error, bool) {
	for {
		if record, err := csvr.Read(); err != nil {
			if err != io.EOF {
				return err, false
			} else {
				break
			}
		} else if len(record) < 1 {
			continue
		} else {
			ln := record[0]
			tk := tokenizer.SplitText(ln)
			pn := getPname(tk)
			desc := getDesc(tk)
			if err = csvw.Write([]string{ln, pn, desc}); err != nil {
				return err, true
			}
		}
	}
	return nil, false
}

func GetPnameFunc(lCamel, uCamel, lSnake, uSnake bool) func([]*Token) string {
	var pnameFunc func([]string) string
	if lCamel {
		pnameFunc = ToLowerCamelCase
	} else if uCamel {
		pnameFunc = ToUpperCamelCase
	} else if lSnake {
		pnameFunc = ToLowerSnakeCase
	} else if uSnake {
		pnameFunc = ToUpperSnakeCase
	} else {
		pnameFunc = func(t []string) string {
			return strings.Join(t, " ")
		}
	}
	return func(token []*Token) string {
		t := make([]string, 0, len(token))
		for _, tk := range token {
			for _, n := range tk.Name {
				t = append(t, n)
			}
		}
		return pnameFunc(t)
	}
}

func GetDescFunc() func([]*Token) string {
	return func(token []*Token) string {
		t := make([]string, 0, len(token))
		for _, tk := range token {
			if tk.OK {
				t = append(t, fmt.Sprintf("%s=>%s", tk.Word, tk.Name))
			} else {
				t = append(t, fmt.Sprintf("%s=*", tk.Word))
			}
		}
		return strings.Join(t, "|")
	}
}
