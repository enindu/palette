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
	Writer     Writer     // Output writer
	Styles     []Style    // Text styles
	Foreground Foreground // Text foreground color
	Background Background // Text background color
}

// Print formats i and a, and writes it to p.Writer. It returns number bytes
// written and any error occurred.
func (p *Printer) Print(i string, a ...any) (int, error) {
	format := p.format(i, a...)
	return fmt.Fprint(p.Writer, format)
}

// SetWriter sets p.Writer using w. It returns a pointer to p.
func (p *Printer) SetWriter(w Writer) *Printer {
	p.Writer = w
	return p
}

// SetStyles sets p.Styles using s. It returns a pointer to p.
func (p *Printer) SetStyles(s ...Style) *Printer {
	p.Styles = s
	return p
}

// SetForeground sets p.Foreground using f. It returns a pointer to p.
func (p *Printer) SetForeground(f Foreground) *Printer {
	p.Foreground = f
	return p
}

// SetBackground sets p.Background using b. It returns a pointer to p.
func (p *Printer) SetBackground(b Background) *Printer {
	p.Background = b
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
	for _, style := range p.Styles {
		styles = fmt.Sprintf("%v%v;", styles, style)
	}
	return fmt.Sprintf("\x1b[%v%v;%vm", styles, p.Foreground, p.Background)
}

func (p *Printer) unset() string {
	return "\x1b[0m"
}

// NewPrinterRegu creates a new [Printer] to print regular mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].Writer: [WrRegular]
//   - [Printer].Styles: [StRegular]
//   - [Printer].Foreground: [FgRegular]
//   - [Printer].Background: [BgRegular]
func NewPrinterRegu() *Printer {
	return NewPrinter(FgRegular, BgRegular, StRegular)
}

// NewPrinterSucc creates a new [Printer] to print success mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].Writer: [WrRegular]
//   - [Printer].Styles: [StBold]
//   - [Printer].Foreground: [FgGreen]
//   - [Printer].Background: [BgRegular]
func NewPrinterSucc() *Printer {
	return NewPrinter(FgGreen, BgRegular, StBold)
}

// NewPrinterInfo creates a new [Printer] to print information mesaages. It
// returns a pointer to [Printer] with pre-defined values.
//
//   - [Printer].Writer: [WrRegular]
//   - [Printer].Styles: [StBold]
//   - [Printer].Foreground: [FgBlue]
//   - [Printer].Background: [BgRegular]
func NewPrinterInfo() *Printer {
	return NewPrinter(FgBlue, BgRegular, StBold)
}

// NewPrinterWarn creates a new [Printer] to print warning mesaages. It returns
// a pointer to [Printer] with pre-defined values.
//
//   - [Printer].Writer: [WrRegular]
//   - [Printer].Styles: [StBold]
//   - [Printer].Foreground: [FgYellow]
//   - [Printer].Background: [BgRegular]
func NewPrinterWarn() *Printer {
	return NewPrinter(FgYellow, BgRegular, StBold)
}

// NewPrinterErro creates a new [Printer] to print error mesaages. It returns a
// pointer to [Printer] with pre-defined values.
//
//   - [Printer].Writer: [WrError]
//   - [Printer].Styles: [StBold]
//   - [Printer].Foreground: [FgRed]
//   - [Printer].Background: [BgRegular]
func NewPrinterErro() *Printer {
	return NewPrinter(FgRed, BgRegular, StBold).SetWriter(WrError)
}

// NewPrinter creates a new [Printer] using f, b, and s. It returns a pointer to
// [Printer] with default values.
//
//   - [Printer].Writer: [WrRegular]
func NewPrinter(f Foreground, b Background, s ...Style) *Printer {
	return &Printer{
		Writer:     WrRegular,
		Styles:     s,
		Foreground: f,
		Background: b,
	}
}
