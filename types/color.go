package types

import "errors"

type uint6 uint8 // shhh...

type Color struct {
	R uint6
	G uint6
	B uint6
}

type PaletteIndex uint8

func (clr Color) IsValid() bool {
	return clr.R <= MAX_UINT6 && clr.G <= MAX_UINT6 && clr.B <= MAX_UINT6
}

var (
	InvalidColor = Color{R: 255, G: 255, B: 255}
)

const MAX_UINT6 uint6 = (1 << 6) - 1

var (
	ErrInvalidUint6 = errors.New("invalid uint6")
	ErrInvalidColor = errors.New("invalid color")
)

func FromRGB(r uint6, g uint6, b uint6) (Color, error) {
	if r > MAX_UINT6 {
		return Color{}, ErrInvalidUint6
	}
	if g > MAX_UINT6 {
		return Color{}, ErrInvalidUint6
	}
	if b > MAX_UINT6 {
		return Color{}, ErrInvalidUint6
	}
	return Color{R: r, G: g, B: b}, nil
}

func FromRGBNoErr(r uint6, g uint6, b uint6) Color {
	var out, err = FromRGB(r, g, b)
	if err != nil {
		panic(err)
	}
	return out
}
