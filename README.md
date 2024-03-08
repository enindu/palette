# Palette

Palette is a simple library to add ANSI colors to terminal output. See [documentation](https://pkg.go.dev/github.com/enindu/palette) for more information.

## Features

- Supports 4 text formats and 32 text colors including foreground and background colors.
- You can set output writer. Writer must implements `palette.Writer`, which is an alias to `io.Writer` interface. Default writer is `palette.WriterRegular`, which is also an alias to `os.Stdout`.
- You can set text length to use `palette.Printer` configurations. Default length is 0, which means entire text will be used by `palette.Printer` configurations.
- You can chain everything in one line using setters.

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
    printer := palette.
		NewPrinter(palette.FormatBold, palette.ForegroundRed, palette.BackgroundRegular).
		SetWriter(palette.WriterError).
		SetLength(4)
	printer.Print("ERRO this is an error message.\n")
}
```
