//go:build run
// +build run

package main

import (
	_ "embed"
	"encoding/json"
	"os"
	"sort"
	"strings"
	"text/template"
)

//go:embed character.json
var jsonCharacters []byte

type Character struct {
	Kana  string `json:"kana"`
	Roman string `json:"roman"`
}

//go:embed contracted.json
var jsonContracted []byte

type Contracted struct {
	Roman1     string `json:"roman1"`
	Roman2     string `json:"roman2"`
	Contracted string `json:"contracted"`
}

var tempStr = `package table

type contracted struct {
	roman1     string
	roman2     string
	contracted string
}

var contractedList = []contracted{
	{{ range .Cont -}}{ roman1: "{{ .Roman1 }}", roman2: "{{ .Roman2 }}", contracted: "{{ .Contracted }}" },
	{{ end -}}
}

var mapCharacters = map[string]string{
	{{ range .Chars -}}"{{ .Kana }}": "{{ .Roman }}",
	{{ end -}}
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
`

func main() {
	var chars []Character
	if err := json.Unmarshal(jsonCharacters, &chars); err != nil {
		panic(err)
	}

	var conts []Contracted
	if err := json.Unmarshal(jsonContracted, &conts); err != nil {
		panic(err)
	}
	sort.Slice(conts, func(i, j int) bool {
		if strings.Compare(conts[i].Roman1, conts[j].Roman1) < 0 {
			return true
		}
		if conts[i].Roman1 == conts[j].Roman1 && strings.Compare(conts[i].Roman2, conts[j].Roman2) < 0 {
			return true
		}
		return false
	})

	var values = struct {
		Chars []Character
		Cont  []Contracted
	}{
		Chars: chars,
		Cont:  conts,
	}

	tpl, err := template.New("jsonCharacters").Parse(tempStr)
	if err != nil {
		panic(err)
	}
	if err := tpl.Execute(os.Stdout, values); err != nil {
		panic(err)
	}
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
