# Palette

Palette is a simple library to add ANSI colors to terminal output. See [documentation](https://pkg.go.dev/github.com/enindu/palette) for more information.

## Features

- Includes 4 text styles and 32 text colors including foreground and background.
- Includes 5 pre-defined printers.
- You can set output writer. Writer must implements `palette.WR`, which is an alias to `io.Writer` interface. Default writer is `palette.WRRegular`, which is also an alias to `os.Stdout`.
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

import "github.com/enindu/palette"

func main() {
	printer := palette.NewPrinter(palette.FGRed, palette.BGRegular, palette.STBold, palette.STUnderline).SetWriter(palette.WRError)
	printer.Print("this is a test text.\n")
}
```
