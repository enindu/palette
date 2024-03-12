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
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Printer represents a new printer. A non-zero printer can print.
type Printer struct {
	writer     io.Writer
	styles     []uint64
	foreground uint64
	background uint64
	buffer     *bytes.Buffer
}

// Print formats i and a, and writes it to p.writer. It returns number bytes
// written and any error occurred.
func (p *Printer) Print(i string, a ...any) (int64, error) {
	defer p.buffer.Reset()
	p.format(i, a...)
	return p.buffer.WriteTo(p.writer)
}

// SetWriter sets p.writer using w. It returns a pointer to p.
func (p *Printer) SetWriter(w io.Writer) *Printer {
	p.writer = w
	return p
}

// SetStyles sets p.styles using s. It returns a pointer to p.
func (p *Printer) SetStyles(s ...uint64) *Printer {
	p.styles = s
	return p
}

// SetForeground sets p.foreground using f. It returns a pointer to p.
func (p *Printer) SetForeground(f uint64) *Printer {
	p.foreground = f
	return p
}

// SetBackground sets p.background using b. It returns a pointer to p.
func (p *Printer) SetBackground(b uint64) *Printer {
	p.background = b
	return p
}

func (p *Printer) format(i string, a ...any) {
	format := fmt.Sprintf(i, a...)
	if !strings.Contains(format, "\n") {
		p.set()
		p.buffer.WriteString(format)
		p.unset()
		return
	}
	lines := strings.SplitAfter(format, "\n")
	length := len(lines)
	if strings.HasSuffix(format, "\n") {
		length = length - 1
	}
	for _, v := range lines[:length] {
		end := ""
		if strings.HasSuffix(v, "\n") {
			v = strings.TrimSuffix(v, "\n")
			end = "\n"
		}
		p.set()
		p.buffer.WriteString(v)
		p.unset()
		p.buffer.WriteString(end)
	}
}

func (p *Printer) set() {
	p.buffer.WriteString("\x1b[")
	for _, v := range p.styles {
		style := strconv.FormatUint(v, 10)
		p.buffer.WriteString(style)
		p.buffer.WriteString(";")
	}
	foreground := strconv.FormatUint(p.foreground, 10)
	background := strconv.FormatUint(p.background, 10)
	p.buffer.WriteString(foreground)
	p.buffer.WriteString(";")
	p.buffer.WriteString(background)
	p.buffer.WriteString("m")
}

func (p *Printer) unset() {
	p.buffer.WriteString("\x1b[0m")
}

// NewPrinterRegu creates a new [Printer] to print regular messages. It returns
// a pointer to [Printer] with pre-defined values.
func NewPrinterRegu() *Printer {
	return NewPrinter(FgRegular, BgRegular, StRegular)
}

// NewPrinterSucc creates a new [Printer] to print success messages. It returns
// a pointer to [Printer] with pre-defined values.
func NewPrinterSucc() *Printer {
	return NewPrinter(FgGreen, BgRegular, StBold)
}

// NewPrinterInfo creates a new [Printer] to print information messages. It
// returns a pointer to [Printer] with pre-defined values.
func NewPrinterInfo() *Printer {
	return NewPrinter(FgBlue, BgRegular, StBold)
}

// NewPrinterWarn creates a new [Printer] to print warning messages. It returns
// a pointer to [Printer] with pre-defined values.
func NewPrinterWarn() *Printer {
	return NewPrinter(FgYellow, BgRegular, StBold)
}

// NewPrinterErro creates a new [Printer] to print error messages. It returns a
// pointer to [Printer] with pre-defined values.
func NewPrinterErro() *Printer {
	return NewPrinter(FgRed, BgRegular, StBold).SetWriter(os.Stderr)
}

// NewPrinter creates a new [Printer] using f, b, and s. It returns a pointer to
// [Printer] with default values.
func NewPrinter(f uint64, b uint64, s ...uint64) *Printer {
	return &Printer{
		writer:     os.Stdout,
		styles:     s,
		foreground: f,
		background: b,
		buffer:     &bytes.Buffer{},
	}
}
