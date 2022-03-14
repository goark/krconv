# [krconv] -- Convert kana-character to roman-alphabet

Convert kana-characters to roman-alphabets (by hepburn romanization)

[![check vulns](https://github.com/goark/krconv/workflows/vulns/badge.svg)](https://github.com/goark/krconv/actions)
[![lint status](https://github.com/goark/krconv/workflows/lint/badge.svg)](https://github.com/goark/krconv/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/goark/krconv/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/goark/krconv.svg)](https://github.com/goark/krconv/releases/latest)

This package is required Go 1.16 or later.

**Migrated repository to [github.com/goark/krconv][krconv]**

## Import

```go
import "github.com/goark/krconv"
```

## Usage

```go
package krconv_test

import (
	"fmt"

	"github.com/goark/krconv"
)

func ExampleConvert() {
	s := "マツエ テッペイ　めっちゃほりでぃ ﾅﾝﾊﾞかげつで まんざい みるんだょっ"
	fmt.Println(krconv.Convert(s))
	//Output:
	//matsue teppei metchahoridei nambakagetsude manzai mirundayotsu
}
```

## Modules Requirement Graph

[![dependency.png](./dependency.png)](./dependency.png)

[krconv]: https://github.com/goark/krconv "goark/krconv: Convert kana-character to roman-alphabet"
