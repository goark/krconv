package table

import (
	"sort"
	"strings"
)

func ExistContractedFirstChar(first string) bool {
	if n := sort.Search(len(contractedList), func(i int) bool { return strings.Compare(contractedList[i].roman1, first) >= 0 }); n < len(contractedList) && contractedList[n].roman1 == first {
		return true
	}
	return false
}

func GetContractedChars(first, second string) string {
	n := sort.Search(len(contractedList), func(i int) bool {
		if strings.Compare(contractedList[i].roman1, first) > 0 {
			return true
		}
		if contractedList[i].roman1 == first && strings.Compare(contractedList[i].roman2, second) >= 0 {
			return true
		}
		return false
	})
	if n < len(contractedList) && contractedList[n].roman1 == first && contractedList[n].roman2 == second {
		return contractedList[n].contracted
	}
	return ""
}

/* Copyright 2021 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
