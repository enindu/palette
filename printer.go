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
	"io"
	"os"
)

// Printer represents a new printer with specified fields. A non-zero value
// printer can be used to write.
type Printer struct {
	writer     io.Writer
	format     Format
	foreground Foreground
	background Background
	length     int
}

// Write writes text to [Printer].writer using input and data. It returns number
// of bytes written and any error encountered.
func (printer *Printer) Write(input string, data ...any) (int, error) {
	text := fmt.Sprintf(input, data...)
	if printer.length > len(text) {
		return 0, fmt.Errorf("write: length is greater than text")
	}
	start := fmt.Sprintf("\x1b[%v;%v;%vm", printer.format, printer.foreground, printer.background)
	end := "\x1b[0;0;0m"
	if printer.length <= 0 {
		return fmt.Fprintf(printer.writer, "%s%s%s", start, text, end)
	}
	return fmt.Fprintf(printer.writer, "%s%s%s%s", start, text[:printer.length], end, text[printer.length:])
}

// SetWriter sets [Printer].writer using writer. It returns a pointer to
// [Printer].
func (printer *Printer) SetWriter(writer io.Writer) *Printer {
	printer.writer = writer
	return printer
}

// SetFormat sets [Printer].format using format. It returns a pointer to
// [Printer].
func (printer *Printer) SetFormat(format Format) *Printer {
	printer.format = format
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

// SetLength sets [Printer].length using length. It returns a pointer to
// [Printer].
//
// length<=0 considered [Printer].length is disabled and entire text will be
// written as in [Printer] configurations.
func (printer *Printer) SetLength(length int) *Printer {
	printer.length = length
	return printer
}

// NewPrinter creates a new [Printer] using format, foreground, and background.
// It return a pointer to [Printer].
func NewPrinter(format Format, foreground Foreground, background Background) *Printer {
	return &Printer{
		writer:     os.Stdout,
		format:     format,
		foreground: foreground,
		background: background,
		length:     0,
	}
}
