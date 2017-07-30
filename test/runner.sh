#!/bin/bash
#
# Copyright 2017 agwlvssainokuni
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

basedir=$(cd $(dirname ${BASH_SOURCE[0]}); pwd)

# (0) ビルドする。
cd ${basedir}
pushd ..; go build; popd

# (1) 物理名生成する。
../go-pname -d dict.tsv -o result.tsv lname.tsv

# (2) 検証する。
echo "BEGIN{diff -u expected.tsv result.tsv}"
diff -u expected.tsv result.tsv
echo "END  {diff -u expected.tsv result.tsv}"

# (3) 後片付けする。
rm -f result.tsv
pushd ..; go clean; popd
