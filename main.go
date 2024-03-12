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
//
// Under the hood, Palette uses [fmt.Sprintf] function to format input, [bytes.Buffer] struct to concatenate strings, and "Writer" method of [io.Writer] interface to write output. It does not use traditional [fmt] functions to write.
//
// See [repository] for more information.
//
// [repository]: https://github.com/enindu/palette
package palette

const (
	StRegular   uint64 = 0 // Reset style
	StBold      uint64 = 1 // Bold style
	StDim       uint64 = 2 // Dim style
	StItalic    uint64 = 3 // Italic style
	StUnderline uint64 = 4 // Underline style
)

const (
	FgRegular   uint64 = 39 // Reset foreground
	FgBlack     uint64 = 30 // Black foreground
	FgRed       uint64 = 31 // Red foreground
	FgGreen     uint64 = 32 // Green foreground
	FgYellow    uint64 = 33 // Yellow foreground
	FgBlue      uint64 = 34 // Blue foreground
	FgMagenta   uint64 = 35 // Magenta foreground
	FgCyan      uint64 = 36 // Cyan foreground
	FgWhite     uint64 = 37 // White foreground
	FgHiBlack   uint64 = 90 // High intensity black foreground
	FgHiRed     uint64 = 91 // High intensity red foreground
	FgHiGreen   uint64 = 92 // High intensity green foreground
	FgHiYellow  uint64 = 93 // High intensity yellow foreground
	FgHiBlue    uint64 = 94 // High intensity blue foreground
	FgHiMagenta uint64 = 95 // High intensity magenta foreground
	FgHiCyan    uint64 = 96 // High intensity cyan foreground
	FgHiWhite   uint64 = 97 // High intensity white foreground
)

const (
	BgRegular   uint64 = 49  // Reset background
	BgBlack     uint64 = 40  // Black background
	BgRed       uint64 = 41  // Red background
	BgGreen     uint64 = 42  // Green background
	BgYellow    uint64 = 43  // Yellow background
	BgBlue      uint64 = 44  // Blue background
	BgMagenta   uint64 = 45  // Magenta background
	BgCyan      uint64 = 46  // Cyan background
	BgWhite     uint64 = 47  // White background
	BgHiBlack   uint64 = 100 // High intensity black background
	BgHiRed     uint64 = 101 // High intensity red background
	BgHiGreen   uint64 = 102 // High intensity green background
	BgHiYellow  uint64 = 103 // High intensity yellow background
	BgHiBlue    uint64 = 104 // High intensity blue background
	BgHiMagenta uint64 = 105 // High intensity magenta background
	BgHiCyan    uint64 = 106 // High intensity cyan background
	BgHiWhite   uint64 = 107 // High intensity white background
)
