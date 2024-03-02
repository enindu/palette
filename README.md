# PALETTE

Palette is a simple library designed to add colors to terminal output. Since Palette utilizes ANSI colors, it will only function on terminals that support ANSI.

## FEATURES

- Supports 4 formats.
    - Bold
    - Dim
    - Italic
    - Underline

- Supports 32 colors.
    - Foreground: black/high intensity black
    - Foreground: red/high intensity red
    - Foreground: green/high intensity green
    - Foreground: yellow/high intensity yellow
    - Foreground: blue/high intensity blue
    - Foreground: magenta/high intensity magenta
    - Foreground: cyan/high intensity cyan
    - Foreground: white/high intensity white
    - Background: black/high intensity black
    - Background: red/high intensity red
    - Background: green/high intensity green
    - Background: yellow/high intensity yellow
    - Background: blue/high intensity blue
    - Background: magenta/high intensity magenta
    - Background: cyan/high intensity cyan
    - Background: white/high intensity white

- You can manually set the output writer. The writer must implement the `io.Writer` interface. The default output writer is `os.Stdout`.

```go
color := palette.
    NewColor(palette.Bold, palette.FGRed, palette.BGRegular).
    SetWriter(os.Stderr)

// This message will use "os.Stderr" as the writer
color.Print("erro this is an error message")
```

- You can set the text length that must be colored. The default length is `0`, which means the full text will be colored.

```go
color := palette.
    NewColor(palette.Bold, palette.FGRed, palette.BGRegular).
    SetLength(4)

// Only "erro" will be colored
color.Print("erro this is an error message")
```

- You can chain everything in one line.

```go
palette.
    NewColor(palette.Bold, palette.FGRed, palette.BGRegular).
    SetWriter(os.Stderr).
    SetLength(4).
    Print("erro this is an error message")
```

## INSTALL

You can add Palette to your application using the `go get` command.

```
go get github.com/enindu/palette
```

## USAGE

Here's an example of how to use Palette.

```go
package main

import (
    "os"

    "github.com/enindu/palette"
)

erro := palette.
    NewColor(palette.Bold, palette.FGRed, palette.BGRegular).
    SetWriter(os.Stderr).
    SetLength(4)

erro.Print("erro this is an error message")
```
