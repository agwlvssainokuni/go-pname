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

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"

	"github.com/agwlvssainokuni/go-pname/pname"
)

func processMain(csvr *csv.Reader, csvw *csv.Writer, tokenizer pname.Tokenizer, getPname func([]*pname.Token) string, getDesc func([]*pname.Token) string) (error, error) {
	for {

		record, err := csvr.Read()
		if err != nil {
			if err != io.EOF {
				return err, nil
			} else {
				break
			}
		}

		if len(record) < 1 {
			continue
		}

		ln := record[0]
		tk := tokenizer.SplitText(ln)
		pn := getPname(tk)
		desc := getDesc(tk)
		err = csvw.Write([]string{ln, pn, desc})
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func getPnameFunc(lCamel, uCamel, lSnake, uSnake bool) func([]*pname.Token) string {
	var pnameFunc func([]string) string
	if lCamel {
		pnameFunc = pname.ToLowerCamelCase
	} else if uCamel {
		pnameFunc = pname.ToUpperCamelCase
	} else if lSnake {
		pnameFunc = pname.ToLowerSnakeCase
	} else if uSnake {
		pnameFunc = pname.ToUpperSnakeCase
	} else {
		pnameFunc = func(t []string) string {
			return strings.Join(t, " ")
		}
	}
	return func(token []*pname.Token) string {
		t := make([]string, 0, len(token))
		for _, tk := range token {
			for _, pn := range tk.Pnm {
				t = append(t, pn)
			}
		}
		return pnameFunc(t)
	}
}

func getDescFunc() func([]*pname.Token) string {
	return func(token []*pname.Token) string {
		t := make([]string, 0, len(token))
		for _, tk := range token {
			if tk.OK {
				t = append(t, fmt.Sprintf("%s=>%s", tk.Lnm, tk.Pnm))
			} else {
				t = append(t, fmt.Sprintf("%s=*", tk.Lnm))
			}
		}
		return strings.Join(t, "|")
	}
}
