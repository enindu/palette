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
	Writer     WR
	Styles     []ST
	Foreground FG
	Background BG
}

// Print prints text to [Printer].writer using input and data. It returns number
// of bytes written and any error encountered.
func (p *Printer) Print(i string, a ...any) (int, error) {
	text := p.text(i, a...)
	return fmt.Fprint(p.Writer, text)
}

// SetWriter sets [Printer].writer using writer. It returns a pointer to
// [Printer].
func (p *Printer) SetWriter(w WR) *Printer {
	p.Writer = w
	return p
}

// SetStyles sets [Printer].styles using styles. It returns a pointer to
// [Printer].
func (p *Printer) SetStyles(s ...ST) *Printer {
	p.Styles = s
	return p
}

// SetForeground sets [Printer].foreground using foreground. It returns a
// pointer to [Printer].
func (p *Printer) SetForeground(f FG) *Printer {
	p.Foreground = f
	return p
}

// SetBackground sets [Printer].background using background. It returns a
// pointer to [Printer].
func (p *Printer) SetBackground(b BG) *Printer {
	p.Background = b
	return p
}

func (p *Printer) text(i string, a ...any) string {
	start := p.start()
	end := p.end()
	text := fmt.Sprintf(i, a...)
	if !strings.Contains(text, "\n") {
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

func (p *Printer) start() string {
	styles := ""
	for _, style := range p.Styles {
		styles = fmt.Sprintf("%v%v;", styles, style)
	}
	return fmt.Sprintf("\x1b[%v%v;%vm", styles, p.Foreground, p.Background)
}

func (p *Printer) end() string {
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
func NewPrinter(f FG, b BG, s ...ST) *Printer {
	return &Printer{
		Writer:     WRRegular,
		Styles:     s,
		Foreground: f,
		Background: b,
	}
}
