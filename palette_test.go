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
	"os"
	"testing"

	fatihcolor "github.com/fatih/color"
	gookitcolor "github.com/gookit/color"
)

func BenchmarkEninduPalette(b *testing.B) {
	out, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		p := NewPrinter(FgRed, BgRegular, StBold).SetWriter(out)
		p.Print("hello world!\n")
	}
}

func BenchmarkFatihColor(b *testing.B) {
	out, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		c := fatihcolor.New(fatihcolor.FgRed, fatihcolor.Bold)
		c.Fprint(out, "hello world!\n")
	}
}

func BenchmarkGookitColor(b *testing.B) {
	out, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		gookitcolor.SetOutput(out)
		c := gookitcolor.New(gookitcolor.Red, gookitcolor.Bold)
		c.Print("hello world!\n")
	}
}
