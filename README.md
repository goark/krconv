# [krconv] -- Convert kana-character to roman-alphabet

Convert kana-characters to roman-alphabets (by hepburn romanization)

[![check vulns](https://github.com/spiegel-im-spiegel/krconv/workflows/vulns/badge.svg)](https://github.com/spiegel-im-spiegel/krconv/actions)
[![lint status](https://github.com/spiegel-im-spiegel/krconv/workflows/lint/badge.svg)](https://github.com/spiegel-im-spiegel/krconv/actions)
[![GitHub license](https://img.shields.io/badge/license-Apache%202-blue.svg)](https://raw.githubusercontent.com/spiegel-im-spiegel/krconv/master/LICENSE)
[![GitHub release](https://img.shields.io/github/release/spiegel-im-spiegel/krconv.svg)](https://github.com/spiegel-im-spiegel/krconv/releases/latest)

This package is required Go 1.16 or later.

## Import

```go
import "github.com/spiegel-im-spiegel/krconv"
```

## Usage

```go
package krconv_test

import (
	"fmt"

	"github.com/spiegel-im-spiegel/krconv"
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

[krconv]: https://github.com/spiegel-im-spiegel/krconv "spiegel-im-spiegel/krconv: Convert kana-character to roman-alphabet"
