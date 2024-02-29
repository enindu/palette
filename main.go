// Palette - Add ANSI colors to terminal output
// Copyright (C) 2024  Enindu Alahapperuma
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, either version 3 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
// PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// this program. If not, see <https://www.gnu.org/licenses/>.

package palette

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	FgRegular   Foreground = 39 // Reset foreground attributes.
	FgBlack     Foreground = 30 // Print black foreground.
	FgRed       Foreground = 31 // Print red foreground.
	FgGreen     Foreground = 32 // Print green foreground.
	FgYellow    Foreground = 33 // Print yellow foreground.
	FgBlue      Foreground = 34 // Print blue foreground.
	FgMagenta   Foreground = 35 // Print magenta foreground.
	FgCyan      Foreground = 36 // Print cyan foreground.
	FgWhite     Foreground = 37 // Print white foreground.
	FgHiBlack   Foreground = 90 // Print high intensity black foreground.
	FgHiRed     Foreground = 91 // Print high intensity red foreground.
	FgHiGreen   Foreground = 92 // Print high intensity green foreground.
	FgHiYellow  Foreground = 93 // Print high intensity yellow foreground.
	FgHiBlue    Foreground = 94 // Print high intensity blue foreground.
	FgHiMagenta Foreground = 95 // Print high intensity magenta foreground.
	FgHiCyan    Foreground = 96 // Print high intensity cyan foreground.
	FgHiWhite   Foreground = 97 // Print high intensity white foreground.
)

const (
	BgRegular   Background = 49  // Reset background attributes.
	BgBlack     Background = 40  // Print black background.
	BgRed       Background = 41  // Print red background.
	BgGreen     Background = 42  // Print green background.
	BgYellow    Background = 43  // Print yellow background.
	BgBlue      Background = 44  // Print blue background.
	BgMagenta   Background = 45  // Print magenta background.
	BgCyan      Background = 46  // Print cyan background.
	BgWhite     Background = 47  // Print white background.
	BgHiBlack   Background = 100 // Print high intensity black background.
	BgHiRed     Background = 101 // Print high intensity red background.
	BgHiGreen   Background = 102 // Print high intensity green background.
	BgHiYellow  Background = 103 // Print high intensity yellow background.
	BgHiBlue    Background = 104 // Print high intensity blue background.
	BgHiMagenta Background = 105 // Print high intensity magenta background.
	BgHiCyan    Background = 106 // Print high intensity cyan background.
	BgHiWhite   Background = 107 // Print high intensity white background.
)

const (
	Regular   Format = 0 // Reset all attributes.
	Bold      Format = 1 // Print bold text.
	Dim       Format = 2 // Print dim text.
	Italic    Format = 3 // Print italic text. This will not work on some terminals.
	Underline Format = 4 // Print underline text. This will not works on some terminals.
)

var errLengthGreaterThanInput error = errors.New("print: length is greater than input")

// Foreground type to define foreground colors.
type Foreground int

// Background type to define background colors.
type Background int

// Format type to define text formats.
type Format int

// Color type to define color structure.
type Color struct {
	writer     io.Writer
	foreground Foreground
	background Background
	format     Format
	length     int
}

// Print input to writer.
func (color *Color) Print(input string, data ...any) (int, error) {
	if color.length > len(input) {
		return 0, errLengthGreaterThanInput
	}
	start, end := color.wrappers(color.foreground, color.background, color.format)
	if color.length <= 0 {
		colorText := fmt.Sprintf(input, data...)
		return fmt.Fprintf(color.writer, "%s%s%s", start, colorText, end)
	}
	colorText := input[:color.length]
	plainText := fmt.Sprintf(input[color.length:], data...)
	return fmt.Fprintf(color.writer, "%s%s%s%s", start, colorText, end, plainText)
}

// Set writer, which must be implemented io.Writer interface.
func (color *Color) SetWriter(writer io.Writer) *Color {
	color.writer = writer
	return color
}

// Set foreground, which must be a type of Foreground.
func (color *Color) SetForeground(foreground Foreground) *Color {
	color.foreground = foreground
	return color
}

// Set background, which must be a type of Background.
func (color *Color) SetBackground(background Background) *Color {
	color.background = background
	return color
}

// Set format, which must be a type of Format.
func (color *Color) SetFormat(format Format) *Color {
	color.format = format
	return color
}

// Set length, which must be a type of int. If length <= 0, it considered length is
// disabled.
func (color *Color) SetLength(length int) *Color {
	color.length = length
	return color
}

func (color *Color) wrappers(foreground Foreground, background Background, format Format) (string, string) {
	return fmt.Sprintf("\x1b[%v;%v;%vm", foreground, background, format), "\x1b[0m"
}

// Create new color structure (writer, foreground, background, format, and length).
//
// Default writer is os.Stdout and default length is 0.
func NewColor(foreground Foreground, background Background, format Format) Color {
	return Color{
		writer:     os.Stdout,
		foreground: foreground,
		background: background,
		format:     format,
		length:     0,
	}
}
