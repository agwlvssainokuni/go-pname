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
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/agwlvssainokuni/go-pname/pname"
)

func main() {

	// 1. コマンドラインオプションを解析する。
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
	flag.StringVar(&paramDictFile, "dict", "dict.csv", "単語辞書ファイルのパス")
	flag.StringVar(&paramDictDelim, "D", " ", "単語辞書の分離文字")
	flag.StringVar(&paramDictDelim, "delim", " ", "単語辞書の分離文字")
	flag.StringVar(&paramOutputFile, "o", "", "結果を出力するファイルのパス")
	flag.StringVar(&paramOutputFile, "output", "", "結果を出力するファイルのパス")
	flag.BoolVar(&paramTsvStdio, "T", false, "標準入出力をTSV形式とする")
	flag.BoolVar(&paramTsvStdio, "tsv", false, "標準入出力をTSV形式とする")
	flag.BoolVar(&paramLowerCamel, "c", false, "camelCase")
	flag.BoolVar(&paramLowerCamel, "lower-camel", false, "camelCase")
	flag.BoolVar(&paramUpperCamel, "C", false, "CamelCase")
	flag.BoolVar(&paramUpperCamel, "upper-camel", false, "CamelCase")
	flag.BoolVar(&paramLowerSnake, "s", false, "snake_case")
	flag.BoolVar(&paramLowerSnake, "lower-snake", false, "snake_case")
	flag.BoolVar(&paramUpperSnake, "S", false, "SNAKE_CASE")
	flag.BoolVar(&paramUpperSnake, "upper-snake", false, "SNAKE_CASE")
	flag.Parse()

	// 2. コマンドラインオプションに応じて要素処理を組み立てる。
	// ・辞書に基づき単語分割する処理。
	var tokenizer pname.Tokenizer
	if dict, err := pname.LoadDictFile(paramDictFile, false, paramDictDelim); err != nil {
		fmt.Fprintf(os.Stderr, "単語辞書ファイル読み込みエラー: %s\n", err.Error())
		os.Exit(-1)
	} else {
		tokenizer = pname.NewDictTokenizer(dict)
	}

	// ・物理名を出力書式に整形する処理。
	getPname := getPnameFunc(paramLowerCamel, paramUpperCamel, paramLowerSnake, paramUpperSnake)
	// ・物理名生成の補足情報を整形する処理。
	getDesc := getDescFunc()

	// ・入力元を形成し主処理を呼び出す、及び、エラーが発生したらメッセージを表示する処理。
	procMain := func(reader io.Reader, isTsv bool, csvw *csv.Writer) bool {
		csvr := csv.NewReader(reader)
		if isTsv {
			csvr.Comma = '\t'
		}
		rderr, wrerr := processMain(csvr, csvw, tokenizer, getPname, getDesc)
		if rderr != nil {
			fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", rderr.Error())
			return false
		}
		if wrerr != nil {
			fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", wrerr.Error())
			return false
		}
		return true
	}

	// 3. 出力先を形成する。
	// ・コマンドラインオプションで指定されれば当該ファイルへ出力する。
	// ・指定されなければ標準出力へ出力する。
	var csvw *csv.Writer
	if len(paramOutputFile) <= 0 {
		csvw = csv.NewWriter(os.Stdout)
		if paramTsvStdio {
			csvw.Comma = '\t'
		}
	} else {
		writer, err := os.Create(paramOutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
			os.Exit(-1)
		}
		defer func() {
			if err := writer.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
			}
		}()
		csvw = csv.NewWriter(writer)
		if strings.HasSuffix(paramOutputFile, ".tsv") {
			csvw.Comma = '\t'
		}
	}

	defer func() {
		csvw.Flush()
		if err := csvw.Error(); err != nil {
			fmt.Fprintf(os.Stderr, "結果出力エラー: %s\n", err.Error())
		}
	}()

	// 4. ファイル入出力と組み合わせて主処理を呼び出す。
	// ・引数が指定されれば当該ファイルを順次開いて主処理を呼び出す。
	// ・指定されなければ標準入力から読み込むものとして主処理を呼び出す。
	if len(flag.Args()) <= 0 {
		ok := procMain(os.Stdin, paramTsvStdio, csvw)
		if !ok {
			os.Exit(-1)
		}
	} else {
		procFile := func(fname string) bool {
			reader, err := os.Open(fname)
			if err != nil {
				fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
				return false
			}
			defer func() {
				if err := reader.Close(); err != nil {
					fmt.Fprintf(os.Stderr, "入力読込エラー: %s\n", err.Error())
				}
			}()
			return procMain(reader, strings.HasSuffix(fname, ".tsv"), csvw)
		}
		for _, fname := range flag.Args() {
			ok := procFile(fname)
			if !ok {
				os.Exit(-1)
			}
		}
	}
}
