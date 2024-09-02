package impl

import "github.com/averseabfun/flux/types"

type FlatSampler struct {
	color types.PaletteIndex
}

func (fs *FlatSampler) GetColor() types.PaletteIndex {
	return fs.color
}

func (fs *FlatSampler) SetColor(color types.PaletteIndex) {
	fs.color = color
}

func (fs *FlatSampler) GetAtPoint(_ types.SamplerPoint) types.PaletteIndex {
	return fs.color
}
