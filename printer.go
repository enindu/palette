// This file is part of Palette.
// Copyright (C) 2024 Enindu Alahapperuma
//
// Palette is free software: you can redistribute it and/or modify it under the
// terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// Palette is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// Palette. If not, see <https://www.gnu.org/licenses/>.

package palette

import (
	"fmt"
	"strings"
)

// Printer represents a new printer with specified fields. A non-zero printer
// can be used to print.
type Printer struct {
	writer     WR
	styles     []ST
	foreground FG
	background BG
}

// Print prints text to [Printer].writer using input and data. It returns number
// of bytes written and any error encountered.
func (printer *Printer) Print(input string, data ...any) (int, error) {
	text := printer.text(input, data...)
	return fmt.Fprint(printer.writer, text)
}

// SetWriter sets [Printer].writer using writer. It returns a pointer to
// [Printer].
func (printer *Printer) SetWriter(writer WR) *Printer {
	printer.writer = writer
	return printer
}

// SetStyles sets [Printer].styles using styles. It returns a pointer to
// [Printer].
func (printer *Printer) SetStyles(styles ...ST) *Printer {
	printer.styles = styles
	return printer
}

// SetForeground sets [Printer].foreground using foreground. It returns a
// pointer to [Printer].
func (printer *Printer) SetForeground(foreground FG) *Printer {
	printer.foreground = foreground
	return printer
}

// SetBackground sets [Printer].background using background. It returns a
// pointer to [Printer].
func (printer *Printer) SetBackground(background BG) *Printer {
	printer.background = background
	return printer
}

func (printer *Printer) text(input string, data ...any) string {
	start := printer.start()
	end := printer.end()
	text := fmt.Sprintf(input, data...)
	if !strings.Contains(input, "\n") {
		return fmt.Sprintf("%s%s%s", start, text, end)
	}
	lines := strings.Split(text, "\n")
	length := len(lines) - 1
	text = ""
	for _, line := range lines[:length] {
		text = fmt.Sprintf("%s%s%s%s\n", text, start, line, end)
	}
	return text
}

func (printer *Printer) start() string {
	styles := ""
	for _, style := range printer.styles {
		styles = fmt.Sprintf("%v%v;", styles, style)
	}
	return fmt.Sprintf("\x1b[%v%v;%vm", styles, printer.foreground, printer.background)
}

func (printer *Printer) end() string {
	return "\x1b[0m"
}

// NewPrinterRegu creates a new [Printer] to print regular mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].styles: [STRegular]
//   - [Printer].foreground: [FGRegular]
//   - [Printer].background: [BGRegular]
func NewPrinterRegu() *Printer {
	return NewPrinter(FGRegular, BGRegular, STRegular)
}

// NewPrinterSucc creates a new [Printer] to print success mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].styles: [STBold]
//   - [Printer].foreground: [FGGreen]
//   - [Printer].background: [BGRegular]
func NewPrinterSucc() *Printer {
	return NewPrinter(FGGreen, BGRegular, STBold)
}

// NewPrinterInfo creates a new [Printer] to print information mesaages. It
// returns a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].styles: [STBold]
//   - [Printer].foreground: [FGBlue]
//   - [Printer].background: [BGRegular]
func NewPrinterInfo() *Printer {
	return NewPrinter(FGBlue, BGRegular, STBold)
}

// NewPrinterWarn creates a new [Printer] to print warning mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].styles: [STBold]
//   - [Printer].foreground: [FGYellow]
//   - [Printer].background: [BGRegular]
func NewPrinterWarn() *Printer {
	return NewPrinter(FGYellow, BGRegular, STBold)
}

// NewPrinterErro creates a new [Printer] to print error mesaages. It returns a
// pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WRError]
//   - [Printer].styles: [STBold]
//   - [Printer].foreground: [FGRed]
//   - [Printer].background: [BGRegular]
func NewPrinterErro() *Printer {
	return NewPrinter(FGRed, BGRegular, STBold).SetWriter(WRError)
}

// NewPrinter creates a new [Printer] using foreground, background, and styles.
// It returns a pointer to [Printer] with default values.
//
//   - [Printer].writer: [WriterRegular]
func NewPrinter(foreground FG, background BG, styles ...ST) *Printer {
	return &Printer{
		writer:     WRRegular,
		styles:     styles,
		foreground: foreground,
		background: background,
	}
}
