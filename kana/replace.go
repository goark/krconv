package kana

import (
	"strings"
	"unicode"
)

var (
	kanaCase = unicode.SpecialCase{
		unicode.CaseRange{Lo: 'ぁ', Hi: 'ゖ', Delta: [unicode.MaxCase]rune{'ァ' - 'ぁ', 0, 0}},
		unicode.CaseRange{Lo: 'ゝ', Hi: 'ゞ', Delta: [unicode.MaxCase]rune{'ヽ' - 'ゝ', 0, 0}},
		unicode.CaseRange{Lo: 'ァ', Hi: 'ヶ', Delta: [unicode.MaxCase]rune{0, 'ぁ' - 'ァ', 0}},
		unicode.CaseRange{Lo: 'ヽ', Hi: 'ヾ', Delta: [unicode.MaxCase]rune{0, 'ゝ' - 'ヽ', 0}},
	}
	replacekanaMap = map[string]string{
		string([]rune{'ヷ'}): string([]rune{'わ', 0x3099}),
		string([]rune{'ヸ'}): string([]rune{'ゐ', 0x3099}),
		string([]rune{'ヹ'}): string([]rune{'ゑ', 0x3099}),
		string([]rune{'ヺ'}): string([]rune{'を', 0x3099}),
	}
)

//ReplaceHiragana replaces hiragana from katrakana (full-width kana kcharacter only).
func ReplaceHiragana(txt string) string {
	ss := []string{}
	for k, v := range replacekanaMap {
		ss = append(ss, k, v)
	}
	return strings.ToLowerSpecial(kanaCase, strings.NewReplacer(ss...).Replace(txt))
}

/* Copyright 2020-2021 Spiegel
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
