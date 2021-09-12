package krconv

import (
	"strings"

	"github.com/rivo/uniseg"
	"github.com/spiegel-im-spiegel/krconv/kana"
	"github.com/spiegel-im-spiegel/krconv/table"
	"github.com/spiegel-im-spiegel/krconv/width"
)

func Convert(s string) string {
	//conversion fullwidth katakana
	kanaTxt := width.ConvertStringFold(kana.ReplaceKatakana(s))

	//get characters list
	cl := []string{}
	gr := uniseg.NewGraphemes(kanaTxt)
	next, ok := nextRomanLetters(gr)
	for ok {
		//convert contracted characters（拗音）
		if table.ExistContractedFirstChar(next) {
			char := next
			next, ok = nextRomanLetters(gr)
			if ok {
				if cc, okok := table.GetContractedChars(char, next); okok {
					cl = append(cl, cc)
					next, ok = nextRomanLetters(gr)
				} else {
					cl = append(cl, char)
				}
			} else {
				cl = append(cl, char)
			}
		} else {
			cl = append(cl, next)
			next, ok = nextRomanLetters(gr)
		}
	}

	//check special characters
	cl2 := cl
	cl = make([]string, 0, len(cl2))
	for i := 0; i < len(cl2); i++ {
		switch cl2[i] {
		case "xya", "xyu", "xyo": //single 拗音（ゃゅょ）
			cl = append(cl, cl2[i][1:]) //case toupper
		case "n": //撥音（ん）
			if i < len(cl2)-1 && (strings.HasPrefix(cl2[i+1], "b") || strings.HasPrefix(cl2[i+1], "m") || strings.HasPrefix(cl2[i+1], "p")) {
				cl = append(cl, "m") //set character 'm'
			} else {
				cl = append(cl, cl2[i])
			}
		case "xtsu": //促音（っ）
			if i >= len(cl2)-1 {
				cl = append(cl, cl2[i][1:]) //case toupper
			} else if strings.HasPrefix(cl2[i+1], "ch") {
				cl = append(cl, "t") //set letter 't'
			} else {
				cl = append(cl, cl2[i+1][:1]) //repeat the first character in next letter
			}
		default:
			cl = append(cl, cl2[i])
		}
	}

	return strings.Join(cl, "")
}

func nextRomanLetters(gr *uniseg.Graphemes) (string, bool) {
	if gr.Next() {
		return table.RomanLetter(gr.Str()), true
	}
	return "", false
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
