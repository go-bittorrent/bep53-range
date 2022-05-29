bep53-range is a library for parsing bep53 ranges.

[![Go Reference](https://pkg.go.dev/badge/github.com/go-bittorrent/bep53-range.svg)](https://pkg.go.dev/github.com/go-bittorrent/bep53-range)
[![x](https://github.com/go-bittorrent/bep53-range/actions/workflows/ci.yml/badge.svg)](https://github.com/go-bittorrent/bep53-range/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/go-bittorrent/bep53-range/branch/main/graph/badge.svg?token=snEWZsLYEg)](https://codecov.io/gh/go-bittorrent/bep53-range)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-bittorrent/bep53-range)](https://goreportcard.com/report/github.com/go-bittorrent/bep53-range)

----

## Installation
```bash
go get github.com/go-bittorrent/bep53-range
```

## Example

```go
package main

import (
	"fmt"

	"github.com/go-bittorrent/bep53-range/bep53"
)

func main() {
	parsed, err := bep53.Parse([]string{"1-5", "2", "6-12"})
	if err != nil {
		panic(err)
	}

	fmt.Println(parsed)                // [1 2 3 4 5 5 6 7 8 9 10 11 12]
	fmt.Println(bep53.Compose(parsed)) // [1-5 2 6-12]
}

```

## License
MIT