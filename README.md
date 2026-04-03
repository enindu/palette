# Palette

Palette is a simple library to add ANSI colors to command line output.

Under the hood, Palette uses [fmt.Sprintf](https://pkg.go.dev/fmt#Sprintf) function to format input, [bytes.Buffer](https://pkg.go.dev/bytes#Buffer) struct to concatenate strings, and [bytes.Buffer.WriteTo](https://pkg.go.dev/bytes#Buffer.WriteTo) method to write output. It does not use traditional [fmt](https://pkg.go.dev/fmt) functions to write.

## Features

- Includes 5 text styles and 32 text colors including foreground and background.
- Includes 5 pre-defined printers.
- You can set output writer. Writer must implements [io.Writer](https://pkg.go.dev/io#Writer) interface. Default writer is [os.Stdout](https://pkg.go.dev/os#pkg-variables).
- You can chain everything in one line using setters.

## Install

You can add Palette using `go get` command.

```shell
go get -u github.com/enindu/palette
```

## Usage

Here's an example of how to use Palette.

```go
package main

import (
	"os"

	"github.com/enindu/palette"
)

func main() {
	p := palette.NewPrinter(palette.FgRed, palette.BgRegular, palette.StBold, palette.StUnderline).SetWriter(os.Stderr)
	p.Print("hello world!\n")
}
```

## Benchmark

Palette is an experimental library. It is not even close to competing with other well-designed libraries. My sole aim was to create a lightweight library with minimal allocations and faster execution for use in hot paths.

```
$ go test -bench . -benchtime 10s -benchmem
goos: linux
goarch: amd64
pkg: github.com/enindu/palette
cpu: AMD Ryzen 5 5600H with Radeon Graphics         
BenchmarkEninduPalette-12       16999292               702.6 ns/op           216 B/op          6 allocs/op
BenchmarkFatihColor-12           6438104                1861 ns/op             240 B/op          9 allocs/op
BenchmarkGookitColor-12         14796056               842.1 ns/op           257 B/op          7 allocs/op
PASS
ok      github.com/enindu/palette       39.911s
```
