// Palette - Add ANSI colors to terminal output
// Copyright (C) 2024 Enindu Alahapperuma
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
	"fmt"
	"io"
	"os"
)

const (
	Regular   format = 0 // Reset all attributes.
	Bold      format = 1 // Print bold text.
	Dim       format = 2 // Print dim text.
	Italic    format = 3 // Print italic text. This will not work on some terminals.
	Underline format = 4 // Print underline text. This will not works on some terminals.
)

const (
	FgRegular   foreground = 39 // Reset foreground attributes.
	FgBlack     foreground = 30 // Print black foreground.
	FgRed       foreground = 31 // Print red foreground.
	FgGreen     foreground = 32 // Print green foreground.
	FgYellow    foreground = 33 // Print yellow foreground.
	FgBlue      foreground = 34 // Print blue foreground.
	FgMagenta   foreground = 35 // Print magenta foreground.
	FgCyan      foreground = 36 // Print cyan foreground.
	FgWhite     foreground = 37 // Print white foreground.
	FgHiBlack   foreground = 90 // Print high intensity black foreground.
	FgHiRed     foreground = 91 // Print high intensity red foreground.
	FgHiGreen   foreground = 92 // Print high intensity green foreground.
	FgHiYellow  foreground = 93 // Print high intensity yellow foreground.
	FgHiBlue    foreground = 94 // Print high intensity blue foreground.
	FgHiMagenta foreground = 95 // Print high intensity magenta foreground.
	FgHiCyan    foreground = 96 // Print high intensity cyan foreground.
	FgHiWhite   foreground = 97 // Print high intensity white foreground.
)

const (
	BgRegular   background = 49  // Reset background attributes.
	BgBlack     background = 40  // Print black background.
	BgRed       background = 41  // Print red background.
	BgGreen     background = 42  // Print green background.
	BgYellow    background = 43  // Print yellow background.
	BgBlue      background = 44  // Print blue background.
	BgMagenta   background = 45  // Print magenta background.
	BgCyan      background = 46  // Print cyan background.
	BgWhite     background = 47  // Print white background.
	BgHiBlack   background = 100 // Print high intensity black background.
	BgHiRed     background = 101 // Print high intensity red background.
	BgHiGreen   background = 102 // Print high intensity green background.
	BgHiYellow  background = 103 // Print high intensity yellow background.
	BgHiBlue    background = 104 // Print high intensity blue background.
	BgHiMagenta background = 105 // Print high intensity magenta background.
	BgHiCyan    background = 106 // Print high intensity cyan background.
	BgHiWhite   background = 107 // Print high intensity white background.
)

type color struct {
	writer     io.Writer
	format     format
	foreground foreground
	background background
	length     int
}

type format int

type foreground int

type background int

// Print input to writer.
func (color *color) Print(input string, data ...any) (int, error) {
	if color.length > len(input) {
		return 0, fmt.Errorf("print: length is greater than input")
	}
	openWrapper := fmt.Sprintf("\x1b[%v;%v;%vm", color.format, color.foreground, color.background)
	closeWrapper := "\x1b[0m"
	if color.length <= 0 {
		colorText := fmt.Sprintf(input, data...)
		return fmt.Fprintf(color.writer, "%s%s%s", openWrapper, colorText, closeWrapper)
	}
	colorText := input[:color.length]
	plainText := fmt.Sprintf(input[color.length:], data...)
	return fmt.Fprintf(color.writer, "%s%s%s%s", openWrapper, colorText, closeWrapper, plainText)
}

// Set writer, which must be implemented io.Writer interface.
func (color *color) SetWriter(writer io.Writer) *color {
	color.writer = writer
	return color
}

// Set format, which must be a type of Format.
func (color *color) SetFormat(format format) *color {
	color.format = format
	return color
}

// Set foreground, which must be a type of Foreground.
func (color *color) SetForeground(foreground foreground) *color {
	color.foreground = foreground
	return color
}

// Set background, which must be a type of Background.
func (color *color) SetBackground(background background) *color {
	color.background = background
	return color
}

// Set length, which must be a type of int. If length <= 0, it considered length is
// disabled.
func (color *color) SetLength(length int) *color {
	color.length = length
	return color
}

// Create new color structure (writer, format, foreground, background, and length).
//
// Default writer is os.Stdout and default length is 0.
func NewColor(format format, foreground foreground, background background) *color {
	return &color{
		writer:     os.Stdout,
		format:     format,
		foreground: foreground,
		background: background,
		length:     0,
	}
}
