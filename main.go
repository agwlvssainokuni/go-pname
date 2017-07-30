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

package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/agwlvssainokuni/go-pname/pname"
)

func main() {

	var (
		paramDictFile   string
		paramOutputFile string
		paramLowerCamel bool
		paramUpperCamel bool
		paramLowerSnake bool
		paramUpperSnake bool
	)

	flag.StringVar(&paramDictFile, "d", "dict.csv", "単語辞書ファイルのパス")
	flag.StringVar(&paramOutputFile, "o", "result.csv", "結果を出力するファイルのパス")
	flag.BoolVar(&paramLowerCamel, "c", false, "camelCase")
	flag.BoolVar(&paramUpperCamel, "C", false, "CamelCase")
	flag.BoolVar(&paramLowerSnake, "s", false, "snake_case")
	flag.BoolVar(&paramUpperSnake, "S", false, "SNAKE_CASE")
	flag.Parse()

	var tokenizer pname.Tokenizer
	if dict, err := pname.LoadDictFile(paramDictFile, false); err != nil {
		fmt.Fprintf(os.Stderr, "単語辞書ファイル読み込みエラー: %s\n", err.Error())
		os.Exit(-1)
	} else {
		tokenizer = pname.NewDictTokenizer(dict)
	}

	if writer, err := os.Create(paramOutputFile); err != nil {
		fmt.Fprintf(os.Stderr, "結果ファイル出力エラー: %s\n", err.Error())
		os.Exit(-1)
	} else {

		csvw := csv.NewWriter(writer)
		if strings.HasSuffix(paramOutputFile, ".tsv") {
			csvw.Comma = '\t'
		}

		for _, fname := range flag.Args() {
			if lines, err := loadInputFile(fname); err != nil {
				fmt.Fprintf(os.Stderr, "入力ファイル読み込みエラー: %s\n", err.Error())
			} else {
				for _, ln := range lines {
					token := tokenizer.SplitText(ln)
					pn := getPname(token, paramLowerCamel, paramUpperCamel, paramLowerSnake, paramUpperSnake)
					desc := getDescription(token)
					if err = csvw.Write([]string{ln, pn, desc}); err != nil {
						fmt.Fprintf(os.Stderr, "結果ファイル出力エラー: %s\n", err.Error())
						goto EXIT
					}
				}
			}
		}

	EXIT:
		csvw.Flush()
		if err := csvw.Error(); err != nil {
			fmt.Fprintf(os.Stderr, "結果ファイル出力エラー: %s\n", err.Error())
			os.Exit(-1)
		}
		if err := writer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "結果ファイル出力エラー: %s\n", err.Error())
			os.Exit(-1)
		}
	}
}

func loadInputFile(fname string) ([]string, error) {
	if reader, err := os.Open(fname); err != nil {
		return nil, err
	} else {
		csvr := csv.NewReader(reader)
		if strings.HasSuffix(fname, ".tsv") {
			csvr.Comma = '\t'
		}
		if records, err := csvr.ReadAll(); err != nil {
			return nil, err
		} else {
			if err = reader.Close(); err != nil {
				return nil, err
			}
			result := make([]string, 0, len(records))
			for _, r := range records {
				if len(r) < 1 {
					continue
				}
				result = append(result, r[0])
			}
			return result, nil
		}
	}
}

func getPname(token []*pname.Token, lCamel, uCamel, lSnake, uSnake bool) string {
	t := make([]string, 0, len(token))
	for _, tk := range token {
		t = append(t, tk.Name)
	}
	if lCamel {
		return pname.ToLowerCamelCase(t)
	} else if uCamel {
		return pname.ToUpperCamelCase(t)
	} else if lSnake {
		return pname.ToLowerSnakeCase(t)
	} else if uSnake {
		return pname.ToUpperSnakeCase(t)
	} else {
		return strings.Join(t, " ")
	}
}

func getDescription(token []*pname.Token) string {
	t := make([]string, 0, len(token))
	for _, tk := range token {
		if tk.OK {
			t = append(t, fmt.Sprintf("%s=>%s", tk.Word, tk.Name))
		} else {
			t = append(t, fmt.Sprintf("%s=*", tk.Word))
		}
	}
	return strings.Join(t, ",")
}
