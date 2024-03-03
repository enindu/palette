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

// Package palette is a simple library to add colors to the terminal output.
// Since this package utilizes ANSI colors, it will works only on ANSI supported
// terminals. See [documentation] for more information.
//
// [documentation]: https://github.com/enindu/palette
package palette

import (
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
	FGRegular   foreground = 39 // Reset foreground attributes.
	FGBlack     foreground = 30 // Print black foreground.
	FGRed       foreground = 31 // Print red foreground.
	FGGreen     foreground = 32 // Print green foreground.
	FGYellow    foreground = 33 // Print yellow foreground.
	FGBlue      foreground = 34 // Print blue foreground.
	FGMagenta   foreground = 35 // Print magenta foreground.
	FGCyan      foreground = 36 // Print cyan foreground.
	FGWhite     foreground = 37 // Print white foreground.
	FGHiBlack   foreground = 90 // Print high intensity black foreground.
	FGHiRed     foreground = 91 // Print high intensity red foreground.
	FGHiGreen   foreground = 92 // Print high intensity green foreground.
	FGHiYellow  foreground = 93 // Print high intensity yellow foreground.
	FGHiBlue    foreground = 94 // Print high intensity blue foreground.
	FGHiMagenta foreground = 95 // Print high intensity magenta foreground.
	FGHiCyan    foreground = 96 // Print high intensity cyan foreground.
	FGHiWhite   foreground = 97 // Print high intensity white foreground.
)

const (
	BGRegular   background = 49  // Reset background attributes.
	BGBlack     background = 40  // Print black background.
	BGRed       background = 41  // Print red background.
	BGGreen     background = 42  // Print green background.
	BGYellow    background = 43  // Print yellow background.
	BGBlue      background = 44  // Print blue background.
	BGMagenta   background = 45  // Print magenta background.
	BGCyan      background = 46  // Print cyan background.
	BGWhite     background = 47  // Print white background.
	BGHiBlack   background = 100 // Print high intensity black background.
	BGHiRed     background = 101 // Print high intensity red background.
	BGHiGreen   background = 102 // Print high intensity green background.
	BGHiYellow  background = 103 // Print high intensity yellow background.
	BGHiBlue    background = 104 // Print high intensity blue background.
	BGHiMagenta background = 105 // Print high intensity magenta background.
	BGHiCyan    background = 106 // Print high intensity cyan background.
	BGHiWhite   background = 107 // Print high intensity white background.
)

// This is a first class function of the palette package. It takes
// format(palette.format), foreground(palette.foreground), and
// background(palette.background) as parameters and returns a pointer to the
// palette.Color struct, which can use to invoke palette.Color struct methods.
//
// This function is used to create and configure a new color, using parameter
// and default values.
func NewColor(format format, foreground foreground, background background) *Color {
	return &Color{
		writer:     os.Stdout,
		format:     format,
		foreground: foreground,
		background: background,
		length:     0,
	}
}
