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
)

type color struct {
	writer     io.Writer
	format     format
	foreground foreground
	background background
	length     int
}

// This is a method of the palette.color struct. It takes input(string) and
// ...any as parameters and returns number of bytes written and an error if any
// error occurs.
//
// This method is used to print the formatted input data to the
// palette.color.writer. If the palette.color.length is greater than the input
// length, this method will returns an error. If the palette.color.length is
// equals or less than zero, which means the palette.color.length is disabled,
// this method will prints the entire formatted text in color.
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

// This is a method of the palette.color struct. It takes writer(io.Writer) as
// parameter and returns a pointer to palette.color struct.
//
// This method is used to update the palette.color.writer field, using parameter
// values.
func (color *color) SetWriter(writer io.Writer) *color {
	color.writer = writer
	return color
}

// This is a method of the palette.color struct. It takes format(palette.format)
// as parameter and returns a pointer to palette.color struct.
//
// This method is used to update the palette.color.format field, using parameter
// values.
func (color *color) SetFormat(format format) *color {
	color.format = format
	return color
}

// This is a method of the palette.color struct. It takes
// foreground(palette.foreground) as parameter and returns a pointer to
// palette.color struct.
//
// This method is used to update the palette.color.foreground field, using
// parameter values.
func (color *color) SetForeground(foreground foreground) *color {
	color.foreground = foreground
	return color
}

// This is a method of the palette.color struct. It takes
// background(palette.background) as parameter and returns a pointer to
// palette.color struct.
//
// This method is used to update the palette.color.background field, using
// parameter values.
func (color *color) SetBackground(background background) *color {
	color.background = background
	return color
}

// This is a method of the palette.color struct. It takes length(int) as
// parameter and returns a pointer to palette.color struct.
//
// This method is used to update the palette.color.length field, using
// parameter values.
func (color *color) SetLength(length int) *color {
	color.length = length
	return color
}
