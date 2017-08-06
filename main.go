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
		paramDictDelim  string
		paramOutputFile string
		paramTsvStdio   bool
		paramLowerCamel bool
		paramUpperCamel bool
		paramLowerSnake bool
		paramUpperSnake bool
	)

	flag.StringVar(&paramDictFile, "d", "dict.csv", "単語辞書ファイルのパス")
	flag.StringVar(&paramDictDelim, "D", " ", "単語辞書の分離文字")
	flag.StringVar(&paramOutputFile, "o", "", "結果を出力するファイルのパス")
	flag.BoolVar(&paramTsvStdio, "T", false, "標準入出力をTSV形式とする")
	flag.BoolVar(&paramLowerCamel, "c", false, "camelCase")
	flag.BoolVar(&paramUpperCamel, "C", false, "CamelCase")
	flag.BoolVar(&paramLowerSnake, "s", false, "snake_case")
	flag.BoolVar(&paramUpperSnake, "S", false, "SNAKE_CASE")
	flag.Parse()

	var tokenizer pname.Tokenizer
	if dict, err := pname.LoadDictFile(paramDictFile, false, paramDictDelim); err != nil {
		fmt.Fprintf(os.Stderr, "単語辞書ファイル読み込みエラー: %s\n", err.Error())
		os.Exit(-1)
	} else {
		tokenizer = pname.NewDictTokenizer(dict)
	}

	getPname := pname.GetPnameFunc(paramLowerCamel, paramUpperCamel, paramLowerSnake, paramUpperSnake)
	getDesc := pname.GetDescFunc()

	var (
		writer, reader *os.File
		csvw           *csv.Writer
		csvr           *csv.Reader
		err            error
	)

	if len(paramOutputFile) <= 0 {

		csvw = csv.NewWriter(os.Stdout)
		if paramTsvStdio {
			csvw.Comma = '\t'
		}
	} else {

		if writer, err = os.Create(paramOutputFile); err != nil {
			fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
			os.Exit(-1)
		}

		csvw = csv.NewWriter(writer)
		if strings.HasSuffix(paramOutputFile, ".tsv") {
			csvw.Comma = '\t'
		}
	}

	if len(flag.Args()) <= 0 {

		csvr = csv.NewReader(os.Stdin)
		if paramTsvStdio {
			csvr.Comma = '\t'
		}

		var wrerr bool
		if err, wrerr = pname.Process(csvr, csvw, tokenizer, getPname, getDesc); err != nil {
			if wrerr {
				fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
			} else {
				fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
			}
		}
	} else {

		for _, fname := range flag.Args() {
			if reader, err = os.Open(fname); err != nil {
				fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
				continue
			}

			csvr = csv.NewReader(reader)
			if strings.HasSuffix(fname, ".tsv") {
				csvr.Comma = '\t'
			}

			var wrerr bool
			if err, wrerr = pname.Process(csvr, csvw, tokenizer, getPname, getDesc); err != nil {
				if wrerr {
					fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
				} else {
					fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
				}
			}

			if err = reader.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
			}

			if wrerr {
				goto EXIT
			}
		}
	}

EXIT:
	csvw.Flush()
	if err = csvw.Error(); err != nil {
		fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
		os.Exit(-1)
	}
	if writer != nil {
		if err = writer.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
			os.Exit(-1)
		}
	}
}
