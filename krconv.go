package krconv

import (
	"strings"

	"github.com/rivo/uniseg"
	"github.com/spiegel-im-spiegel/krconv/kana"
	"github.com/spiegel-im-spiegel/krconv/table"
)

func Convert(s string) string {
	kanaTxt := kana.ConvertStringHiragana(s)

	//get character list
	cl := []string{}
	gr := uniseg.NewGraphemes(kanaTxt)
	for gr.Next() {
		cl = append(cl, table.RomanLetters(gr.Str()))
	}
	cl2 := make([]string, 0, len(cl))
	//convert contracted character
	for i := 0; i < len(cl); i++ {
		if table.ExistContractedFirstChar(cl[i]) && i < len(cl)-1 {
			if cc := table.GetContractedChars(cl[i], cl[i+1]); len(cc) > 0 {
				cl2 = append(cl2, cc)
				i++
			} else {
				cl2 = append(cl2, cl[i])
			}
		} else {
			cl2 = append(cl2, cl[i])
		}
	}
	cl = make([]string, 0, len(cl2))
	for i := 0; i < len(cl2); i++ {
		switch cl2[i] {
		case "xya", "xyu", "xyo": //単独拗音（ゃゅょ） to upper case
			cl = append(cl, cl2[i][1:])
		case "n": //撥音（ん）
			if i < len(cl2)-1 && (strings.HasPrefix(cl2[i+1], "b") || strings.HasPrefix(cl2[i+1], "m") || strings.HasPrefix(cl2[i+1], "p")) {
				cl = append(cl, "m")
			} else {
				cl = append(cl, cl2[i])
			}
		case "xtsu": //促音（っ）
			if i >= len(cl2)-1 {
				cl = append(cl, cl2[i][1:]) //促音 to upper case
			} else if strings.HasPrefix(cl2[i+1], "ch") {
				cl = append(cl, "t")
			} else {
				cl = append(cl, cl2[i+1][:1])
			}
		default:
			cl = append(cl, cl2[i])
		}
	}

	return strings.Join(cl, "")
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
