# Unofficial cpubenchmark.net dataset utility

[![Go Reference](https://img.shields.io/badge/go-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/github.com/elliotwutingfeng/cpubenchmarknet)
[![Go Report Card](https://goreportcard.com/badge/github.com/elliotwutingfeng/cpubenchmarknet?style=for-the-badge)](https://goreportcard.com/report/github.com/elliotwutingfeng/cpubenchmarknet)
[![Coveralls](https://img.shields.io/coverallsCoverage/github/elliotwutingfeng/cpubenchmarknet?logo=coveralls&style=for-the-badge)](https://coveralls.io/github/elliotwutingfeng/cpubenchmarknet?branch=main)
[![License](https://img.shields.io/badge/LICENSE-MIT-GREEN?style=for-the-badge)](LICENSE)

An unofficial library for downloading the [CPU Mega List](https://cpubenchmark.net/CPU_mega_page.html) dataset by [PassMark Software](https://passmark.com).

Spot any bugs? [Report them](https://github.com/elliotwutingfeng/cpubenchmarknet/issues).

**Disclaimer:** _This project is not sponsored, endorsed, or otherwise affiliated with PassMark Software._

## Requirements

- Go 1.20

## Basic Example

The following snippet downloads the dataset as a JSON string and prints it to stdout.

```go
package main

import (
    "fmt"

    "github.com/elliotwutingfeng/cpubenchmarknet"
)

func main() {
    CPUMegaList, err := cpubenchmarknet.GetCPUMegaList()
    if err == nil {
        fmt.Println(CPUMegaList) // JSON string
    } else {
        fmt.Println(err)
    }
}
```

## References

- [PassMark Software Terms of Use](https://passmark.com/legal/disclaimer.php)
