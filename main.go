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
	STRegular   ST = 0 // Reset style
	STBold      ST = 1 // Bold style
	STDim       ST = 2 // Dim style
	STItalic    ST = 3 // Italic style
	STUnderline ST = 4 // Underline style
)

const (
	FGRegular   FG = 39 // Reset foreground
	FGBlack     FG = 30 // Black foreground
	FGRed       FG = 31 // Red foreground
	FGGreen     FG = 32 // Green foreground
	FGYellow    FG = 33 // Yellow foreground
	FGBlue      FG = 34 // Blue foreground
	FGMagenta   FG = 35 // Magenta foreground
	FGCyan      FG = 36 // Cyan foreground
	FGWhite     FG = 37 // White foreground
	FGHiBlack   FG = 90 // High intensity black foreground
	FGHiRed     FG = 91 // High intensity red foreground
	FGHiGreen   FG = 92 // High intensity green foreground
	FGHiYellow  FG = 93 // High intensity yellow foreground
	FGHiBlue    FG = 94 // High intensity blue foreground
	FGHiMagenta FG = 95 // High intensity magenta foreground
	FGHiCyan    FG = 96 // High intensity cyan foreground
	FGHiWhite   FG = 97 // High intensity white foreground
)

const (
	BGRegular   BG = 49  // Reset background
	BGBlack     BG = 40  // Black background
	BGRed       BG = 41  // Red background
	BGGreen     BG = 42  // Green background
	BGYellow    BG = 43  // Yellow background
	BGBlue      BG = 44  // Blue background
	BGMagenta   BG = 45  // Magenta background
	BGCyan      BG = 46  // Cyan background
	BGWhite     BG = 47  // White background
	BGHiBlack   BG = 100 // High intensity black background
	BGHiRed     BG = 101 // High intensity red background
	BGHiGreen   BG = 102 // High intensity green background
	BGHiYellow  BG = 103 // High intensity yellow background
	BGHiBlue    BG = 104 // High intensity blue background
	BGHiMagenta BG = 105 // High intensity magenta background
	BGHiCyan    BG = 106 // High intensity cyan background
	BGHiWhite   BG = 107 // High intensity white background
)

var (
	WRError   WR = os.Stderr // Error writer
	WRRegular WR = os.Stdout // Regular writer
)

// WR represents an output writer.
type WR io.Writer

// ST represents a text style.
type ST int

// FG represents a text foreground color.
type FG int

// BG represents a text background color.
type BG int
