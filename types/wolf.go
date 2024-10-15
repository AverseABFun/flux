package types

type Side uint8
const (
	SideTop = Side(iota)
	SideLeft
	SideRight
	SideBottom
)

type RectWolf struct {
	Start Point // Top-down view
	End Point   // Top-down view
	Color PaletteIndex
	
	ID ObjectID
	World *WorldWolf
}

type WorldWolf struct {
	Objects map[ObjectID]*RectWolf
}