// This file is part of Palette.
// Copyright (C) 2024 Enindu Alahapperuma
//
// Palette is free software: you can redistribute it and/or modify it under the terms
// of the GNU General Public License as published by the Free Software Foundation,
// either version 3 of the License, or (at your option) any later version.
//
// Palette is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
// PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with
// Palette. If not, see <https://www.gnu.org/licenses/>.

package palette

import (
	"fmt"
	"io"
)

type color struct {
	writer     io.Writer
	format     format
	foreground foreground
	background background
	length     int
}

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
