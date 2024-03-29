package classicspace

import (
	"image/color"
)

var (
	Black      = color.NRGBA{0x1B, 0x2A, 0x34,  0xFF}
	Blue       = color.NRGBA{0x1E, 0x5A, 0xA8,  0xFF}
	Encarnado  = color.NRGBA{0xFD, 0x03, 0x04,  0xFF}
	Green      = color.NRGBA{0x00, 0x85, 0x2B,  0xFF}
	Husk       = color.NRGBA{0xB1, 0x94, 0x48,  0xFF}
	LightGray  = color.NRGBA{0x8A, 0x92, 0x8D,  0xFF}
	Meadowlark = color.NRGBA{0xEA, 0xDB, 0x40,  0xFF}
	Orange     = color.NRGBA{0xD6, 0x79, 0x23,  0xFF}
	Red        = color.NRGBA{0xB4, 0x00, 0x00,  0xFF}
	Yellow     = color.NRGBA{0xFA, 0xC8, 0x0A,  0xFF}
	White      = color.NRGBA{0xF4, 0xF4, 0xF4,  0xFF}
)

var Palette color.Palette  = color.Palette{
	Black,
	Blue,
	Encarnado,
	Green,
	Husk,
	LightGray,
	Meadowlark,
	Orange,
	Red,
	Yellow,
	White,
}
