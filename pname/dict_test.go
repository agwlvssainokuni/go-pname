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
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDataCsv_0(t *testing.T) {
	if f, err := dictFilepath("dict_0.csv"); err != nil {
		t.Errorf("failed to dictFilepath() %s", err.Error())
		return
	} else if dict, err := LoadDictFile(f, false, " "); err != nil {
		t.Errorf("failed to LoadDictCsv() %s", err.Error())
		return
	} else {
		assertInt(t, 5, len(dict))
		assertString(t, "A", dict["a"][0])
		assertString(t, "AA", dict["aa"][0])
		assertString(t, "B", dict["b"][0])
		assertString(t, "BB", dict["bb"][0])
		assertString(t, "C0", dict["cc"][0])
		assertString(t, "C1", dict["cc"][1])
	}
}

func TestLoadDataCsv_1(t *testing.T) {
	if f, err := dictFilepath("dict_1.csv"); err != nil {
		t.Errorf("failed to dictFilepath() %s", err.Error())
		return
	} else if dict, err := LoadDictFile(f, true, " "); err != nil {
		t.Errorf("failed to LoadDictCsv() %s", err.Error())
		return
	} else {
		assertInt(t, 5, len(dict))
		assertString(t, "A", dict["a"][0])
		assertString(t, "AA", dict["aa"][0])
		assertString(t, "B", dict["b"][0])
		assertString(t, "BB", dict["bb"][0])
		assertString(t, "C0", dict["cc"][0])
		assertString(t, "C1", dict["cc"][1])
	}
}

func TestLoadDataTsv_0(t *testing.T) {
	if f, err := dictFilepath("dict_0.tsv"); err != nil {
		t.Errorf("failed to dictFilepath() %s", err.Error())
		return
	} else if dict, err := LoadDictFile(f, false, " "); err != nil {
		t.Errorf("failed to LoadDictTsv() %s", err.Error())
		return
	} else {
		assertInt(t, 5, len(dict))
		assertString(t, "A", dict["a"][0])
		assertString(t, "AA", dict["aa"][0])
		assertString(t, "B", dict["b"][0])
		assertString(t, "BB", dict["bb"][0])
		assertString(t, "C0", dict["cc"][0])
		assertString(t, "C1", dict["cc"][1])
	}
}

func TestLoadDataTsv_1(t *testing.T) {
	if f, err := dictFilepath("dict_1.tsv"); err != nil {
		t.Errorf("failed to dictFilepath() %s", err.Error())
		return
	} else if dict, err := LoadDictFile(f, true, " "); err != nil {
		t.Errorf("failed to LoadDictTsv() %s", err.Error())
		return
	} else {
		assertInt(t, 5, len(dict))
		assertString(t, "A", dict["a"][0])
		assertString(t, "AA", dict["aa"][0])
		assertString(t, "B", dict["b"][0])
		assertString(t, "BB", dict["bb"][0])
		assertString(t, "C0", dict["cc"][0])
		assertString(t, "C1", dict["cc"][1])
	}
}

func dictFilepath(fname string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, "test", "dict", fname), nil
}
