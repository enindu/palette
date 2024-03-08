# Palette

Palette is a simple library to add ANSI colors to terminal output. See [documentation](https://pkg.go.dev/github.com/enindu/palette) for more information.

## Features

- Supports 4 text formats and 32 text colors including foreground and background colors.
- You can set output writer. Writer must implements `palette.Writer`, which is an alias to `io.Writer` interface. Default writer is `palette.WriterRegular`, which is also an alias to `os.Stdout`.
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
    formats := []palette.Format{
		palette.FormatBold,
		palette.FormatUnderline,
	}
	printer := palette.NewPrinter(formats, palette.ForegroundRed, palette.BackgroundRegular).SetWriter(palette.WriterError)
	printer.Print("this is a test text.\n")
}
```
