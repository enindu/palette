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
	StRegular   Style = 0 // Reset style
	StBold      Style = 1 // Bold style
	StDim       Style = 2 // Dim style
	StItalic    Style = 3 // Italic style
	StUnderline Style = 4 // Underline style
)

const (
	FgRegular   Foreground = 39 // Reset foreground
	FgBlack     Foreground = 30 // Black foreground
	FgRed       Foreground = 31 // Red foreground
	FgGreen     Foreground = 32 // Green foreground
	FgYellow    Foreground = 33 // Yellow foreground
	FgBlue      Foreground = 34 // Blue foreground
	FgMagenta   Foreground = 35 // Magenta foreground
	FgCyan      Foreground = 36 // Cyan foreground
	FgWhite     Foreground = 37 // White foreground
	FgHiBlack   Foreground = 90 // High intensity black foreground
	FgHiRed     Foreground = 91 // High intensity red foreground
	FgHiGreen   Foreground = 92 // High intensity green foreground
	FgHiYellow  Foreground = 93 // High intensity yellow foreground
	FgHiBlue    Foreground = 94 // High intensity blue foreground
	FgHiMagenta Foreground = 95 // High intensity magenta foreground
	FgHiCyan    Foreground = 96 // High intensity cyan foreground
	FgHiWhite   Foreground = 97 // High intensity white foreground
)

const (
	BgRegular   Background = 49  // Reset background
	BgBlack     Background = 40  // Black background
	BgRed       Background = 41  // Red background
	BgGreen     Background = 42  // Green background
	BgYellow    Background = 43  // Yellow background
	BgBlue      Background = 44  // Blue background
	BgMagenta   Background = 45  // Magenta background
	BgCyan      Background = 46  // Cyan background
	BgWhite     Background = 47  // White background
	BgHiBlack   Background = 100 // High intensity black background
	BgHiRed     Background = 101 // High intensity red background
	BgHiGreen   Background = 102 // High intensity green background
	BgHiYellow  Background = 103 // High intensity yellow background
	BgHiBlue    Background = 104 // High intensity blue background
	BgHiMagenta Background = 105 // High intensity magenta background
	BgHiCyan    Background = 106 // High intensity cyan background
	BgHiWhite   Background = 107 // High intensity white background
)

var (
	WrError   Writer = os.Stderr // Error writer
	WrRegular Writer = os.Stdout // Regular writer
)

// Writer represents an output writer.
type Writer io.Writer

// Style represents a text style.
type Style int

// Foreground represents a text foreground color.
type Foreground int

// Background represents a text background color.
type Background int
