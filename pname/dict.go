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
	"io"
	"os"
	"strings"
)

func LoadDictFile(fname string, withHeader bool) (map[string]string, error) {
	if reader, err := os.Open(fname); err != nil {
		return nil, err
	} else {
		var dict map[string]string
		if dict, err = LoadDict(reader, withHeader, strings.HasSuffix(fname, ".tsv")); err != nil {
			return nil, err
		}
		if err = reader.Close(); err != nil {
			return nil, err
		}
		return dict, nil
	}
}

func LoadDict(r io.Reader, withHeader bool, tsv bool) (map[string]string, error) {
	result := make(map[string]string)
	csvr := csv.NewReader(r)
	if tsv {
		csvr.Comma = '\t'
	}
	for i := 0; ; i++ {
		record, err := csvr.Read()
		if err == io.EOF {
			return result, nil
		} else if err != nil {
			return nil, err
		} else {
			if withHeader && i == 0 {
				continue
			}
			if len(record) < 2 {
				continue
			}
			result[record[0]] = record[1]
		}
	}
}
