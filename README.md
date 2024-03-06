# Palette

Palette is a simple library to add ANSI colors to terminal output. See [documentation](https://pkg.go.dev/github.com/enindu/palette) for more information.

## Features

- Supports 4 formats and 32 colors.
- You can manually set output writer. Writer must implements `io.Writer` interface. Default writer is `os.Stdout`.
- You can manually set text length to written as in `Printer` configurations. Default length is 0, which means entire text will be written as in `Printer` configurations.
- You can chain everything in one line.

See usage for more information.

## Install

You can add Palette to your application using the `go get` command.

```shell
go get github.com/enindu/palette
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
    erro := palette.
        NewPrinter(palette.FormatBold, palette.ForegroundRed, palette.BackgroundRegular).
        SetWriter(os.Stderr).
        SetLength(4)
    erro.Write("ERRO this is an error message.\n")
}
```
