package impl

import (
	"os"
	"strconv"
	"strings"

	"github.com/averseabfun/flux/types"
)

func ImportWolfWorld(path string) (types.WorldWolf, error) {
	var byteData, err = os.ReadFile(path)
	if err != nil {
		return types.WorldWolf{}, err
	}
	var data = string(byteData)
	var out = types.WorldWolf{Objects: make(map[types.ObjectID]*types.RectWolf)}
	for _, d := range strings.Split(data, "\n") {
		var d2 = strings.Split(d, ",")
		objId, err := strconv.Atoi(d2[0])
		if err != nil {
			return out, err
		}
		clr, err := strconv.Atoi(d2[1])
		if err != nil {
			return out, err
		}
		x1, err := strconv.Atoi(d2[2])
		if err != nil {
			return out, err
		}
		y1, err := strconv.Atoi(d2[3])
		if err != nil {
			return out, err
		}
		x2, err := strconv.Atoi(d2[4])
		if err != nil {
			return out, err
		}
		y2, err := strconv.Atoi(d2[5])
		if err != nil {
			return out, err
		}
		out.Objects[types.ObjectID(objId)] = &types.RectWolf{Start: types.Point{X: uint32(x1), Y: uint32(y1)}, End: types.Point{X: uint32(x2), Y: uint32(y2)}, Color: types.PaletteIndex(clr), ID: types.ObjectID(objId), World: &out}
	}
	return out, nil
}