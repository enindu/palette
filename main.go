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

// Package palette is a simple library to add ANSI colors to terminal output.
// See [repository] for more information.
//
// [repository]: https://github.com/enindu/palette
package palette

import (
	"io"
	"os"
)

const (
	FormatRegular   Format = 0 // Reset format
	FormatBold      Format = 1 // Bold format
	FormatDim       Format = 2 // Dim format
	FormatItalic    Format = 3 // Italic format
	FormatUnderline Format = 4 // Underline format
)

const (
	ForegroundRegular   Foreground = 39 // Reset foreground
	ForegroundBlack     Foreground = 30 // Black foreground
	ForegroundRed       Foreground = 31 // Red foreground
	ForegroundGreen     Foreground = 32 // Green foreground
	ForegroundYellow    Foreground = 33 // Yellow foreground
	ForegroundBlue      Foreground = 34 // Blue foreground
	ForegroundMagenta   Foreground = 35 // Magenta foreground
	ForegroundCyan      Foreground = 36 // Cyan foreground
	ForegroundWhite     Foreground = 37 // White foreground
	ForegroundHiBlack   Foreground = 90 // High intensity black foreground
	ForegroundHiRed     Foreground = 91 // High intensity red foreground
	ForegroundHiGreen   Foreground = 92 // High intensity green foreground
	ForegroundHiYellow  Foreground = 93 // High intensity yellow foreground
	ForegroundHiBlue    Foreground = 94 // High intensity blue foreground
	ForegroundHiMagenta Foreground = 95 // High intensity magenta foreground
	ForegroundHiCyan    Foreground = 96 // High intensity cyan foreground
	ForegroundHiWhite   Foreground = 97 // High intensity white foreground
)

const (
	BackgroundRegular   Background = 49  // Reset background
	BackgroundBlack     Background = 40  // Black background
	BackgroundRed       Background = 41  // Red background
	BackgroundGreen     Background = 42  // Green background
	BackgroundYellow    Background = 43  // Yellow background
	BackgroundBlue      Background = 44  // Blue background
	BackgroundMagenta   Background = 45  // Magenta background
	BackgroundCyan      Background = 46  // Cyan background
	BackgroundWhite     Background = 47  // White background
	BackgroundHiBlack   Background = 100 // High intensity black background
	BackgroundHiRed     Background = 101 // High intensity red background
	BackgroundHiGreen   Background = 102 // High intensity green background
	BackgroundHiYellow  Background = 103 // High intensity yellow background
	BackgroundHiBlue    Background = 104 // High intensity blue background
	BackgroundHiMagenta Background = 105 // High intensity magenta background
	BackgroundHiCyan    Background = 106 // High intensity cyan background
	BackgroundHiWhite   Background = 107 // High intensity white background
)

var (
	WriterError   Writer = os.Stderr // Error writer
	WriterRegular Writer = os.Stdout // Regular writer
)

// Writer represents an output writer.
type Writer io.Writer

// Format represents a text format.
type Format int

// Foreground represents a text foreground color.
type Foreground int

// Background represents a text background color.
type Background int
