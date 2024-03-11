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

// Printer represents a new printer. A non-zero printer can print.
type Printer struct {
	writer     Writer
	styles     []Style
	foreground Foreground
	background Background
}

// Print formats i and a, and writes it to p.writer. It returns number bytes
// written and any error occurred.
func (p *Printer) Print(i string, a ...any) (int, error) {
	format := p.format(i, a...)
	return fmt.Fprint(p.writer, format)
}

// SetWriter sets p.writer using w. It returns a pointer to p.
func (p *Printer) SetWriter(w Writer) *Printer {
	p.writer = w
	return p
}

// SetStyles sets p.styles using s. It returns a pointer to p.
func (p *Printer) SetStyles(s ...Style) *Printer {
	p.styles = s
	return p
}

// SetForeground sets p.foreground using f. It returns a pointer to p.
func (p *Printer) SetForeground(f Foreground) *Printer {
	p.foreground = f
	return p
}

// SetBackground sets p.background using b. It returns a pointer to p.
func (p *Printer) SetBackground(b Background) *Printer {
	p.background = b
	return p
}

func (p *Printer) format(i string, a ...any) string {
	set := p.set()
	unset := p.unset()
	format := fmt.Sprintf(i, a...)
	if !strings.Contains(format, "\n") {
		return fmt.Sprintf("%s%s%s", set, format, unset)
	}
	lines := strings.SplitAfter(format, "\n")
	length := len(lines)
	if strings.HasSuffix(format, "\n") {
		length = length - 1
	}
	format = ""
	for _, line := range lines[:length] {
		format = fmt.Sprintf("%s%s%s%s", format, set, line, unset)
	}
	return format
}

func (p *Printer) set() string {
	styles := ""
	for _, style := range p.styles {
		styles = fmt.Sprintf("%v%v;", styles, style)
	}
	return fmt.Sprintf("\x1b[%v%v;%vm", styles, p.foreground, p.background)
}

func (p *Printer) unset() string {
	return "\x1b[0m"
}

// NewPrinterRegu creates a new [Printer] to print regular messages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WrRegular]
//   - [Printer].styles: [StRegular]
//   - [Printer].foreground: [FgRegular]
//   - [Printer].background: [BgRegular]
func NewPrinterRegu() *Printer {
	return NewPrinter(FgRegular, BgRegular, StRegular)
}

// NewPrinterSucc creates a new [Printer] to print success messages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WrRegular]
//   - [Printer].styles: [StBold]
//   - [Printer].foreground: [FgGreen]
//   - [Printer].background: [BgRegular]
func NewPrinterSucc() *Printer {
	return NewPrinter(FgGreen, BgRegular, StBold)
}

// NewPrinterInfo creates a new [Printer] to print information messages. It
// returns a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WrRegular]
//   - [Printer].styles: [StBold]
//   - [Printer].foreground: [FgBlue]
//   - [Printer].background: [BgRegular]
func NewPrinterInfo() *Printer {
	return NewPrinter(FgBlue, BgRegular, StBold)
}

// NewPrinterWarn creates a new [Printer] to print warning messages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WrRegular]
//   - [Printer].styles: [StBold]
//   - [Printer].foreground: [FgYellow]
//   - [Printer].background: [BgRegular]
func NewPrinterWarn() *Printer {
	return NewPrinter(FgYellow, BgRegular, StBold)
}

// NewPrinterErro creates a new [Printer] to print error messages. It returns a
// pointer to [Printer] with pre-defined values.
//
//   - [Printer].writer: [WrError]
//   - [Printer].styles: [StBold]
//   - [Printer].foreground: [FgRed]
//   - [Printer].background: [BgRegular]
func NewPrinterErro() *Printer {
	return NewPrinter(FgRed, BgRegular, StBold).SetWriter(WrError)
}

// NewPrinter creates a new [Printer] using f, b, and s. It returns a pointer to
// [Printer] with default values.
//
//   - [Printer].writer: [WrRegular]
func NewPrinter(f Foreground, b Background, s ...Style) *Printer {
	return &Printer{
		writer:     WrRegular,
		styles:     s,
		foreground: f,
		background: b,
	}
}
