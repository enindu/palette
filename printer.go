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
	writer     Writer
	formats    []Format
	foreground Foreground
	background Background
}

// Print prints text to [Printer].writer using input and data. It returns number
// of bytes written and any error encountered.
func (printer *Printer) Print(input string, data ...any) (int, error) {
	text := printer.text(input, data...)
	return fmt.Fprint(printer.writer, text)
}

// SetWriter sets [Printer].writer using writer. It returns a pointer to
// [Printer].
func (printer *Printer) SetWriter(writer Writer) *Printer {
	printer.writer = writer
	return printer
}

// SetFormats sets [Printer].formats using formats. It returns a pointer to
// [Printer].
func (printer *Printer) SetFormats(formats []Format) *Printer {
	printer.formats = formats
	return printer
}

// SetForeground sets [Printer].foreground using foreground. It returns a
// pointer to [Printer].
func (printer *Printer) SetForeground(foreground Foreground) *Printer {
	printer.foreground = foreground
	return printer
}

// SetBackground sets [Printer].background using background. It returns a
// pointer to [Printer].
func (printer *Printer) SetBackground(background Background) *Printer {
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
	formats := ""
	for _, format := range printer.formats {
		formats = fmt.Sprintf("%v%v;", formats, format)
	}
	return fmt.Sprintf("\x1b[%v%v;%vm", formats, printer.foreground, printer.background)
}

func (printer *Printer) end() string {
	return "\x1b[0m"
}

// NewPrinterSucc creates a new [Printer] to print success mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].formats: [FormatBold]
//   - [Printer].foreground: [ForegroundGreen]
//   - [Printer].background: [BackgroundRegular]
func NewPrinterSucc() *Printer {
	formats := []Format{
		FormatBold,
	}
	return NewPrinter(formats, ForegroundGreen, BackgroundRegular)
}

// NewPrinterInfo creates a new [Printer] to print information mesaages. It
// returns a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].formats: [FormatBold]
//   - [Printer].foreground: [ForegroundBlue]
//   - [Printer].background: [BackgroundRegular]
func NewPrinterInfo() *Printer {
	formats := []Format{
		FormatBold,
	}
	return NewPrinter(formats, ForegroundBlue, BackgroundRegular)
}

// NewPrinterWarn creates a new [Printer] to print warning mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterRegular]
//   - [Printer].formats: [FormatBold]
//   - [Printer].foreground: [ForegroundYellow]
//   - [Printer].background: [BackgroundRegular]
func NewPrinterWarn() *Printer {
	formats := []Format{
		FormatBold,
	}
	return NewPrinter(formats, ForegroundYellow, BackgroundRegular)
}

// NewPrinterErro creates a new [Printer] to print error mesaages. It returns a
// pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WriterError]
//   - [Printer].formats: [FormatBold]
//   - [Printer].foreground: [ForegroundRed]
//   - [Printer].background: [BackgroundRegular]
func NewPrinterErro() *Printer {
	formats := []Format{
		FormatBold,
	}
	return NewPrinter(formats, ForegroundRed, BackgroundRegular).SetWriter(WriterError)
}

// NewPrinter creates a new [Printer] using formats, foreground, and background.
// It returns a pointer to [Printer] with default values.
//
//   - [Printer].writer: [WriterRegular]
func NewPrinter(formats []Format, foreground Foreground, background Background) *Printer {
	return &Printer{
		writer:     WriterRegular,
		formats:    formats,
		foreground: foreground,
		background: background,
	}
}
