package types

import "math"

type Point3D struct {
	X float64
	Y float64
	Z float64
}

type Degree float64
type Radian float64

func (d Degree) ToRadians() Radian {
	return Radian(d * (math.Pi/180))
}

type Object3D struct {
	ID    ObjectID
	World *World3D
}

type Rotation3D struct {
	X Degree
	Y Degree
	Z Degree
}

type World3D struct {
	Objects map[ObjectID]*Object3D
}

type ObjectID uint64

type Collision3D struct {
	At      Point3D
	Object1 ObjectID
	Object2 ObjectID
}

type Ray3D struct {
	Origin   Point3D
	Rotation Rotation3D
	Object3D `collides:"false"`
}

type Poly3D struct {
	Points        []Point3D
	SamplerPoints map[Point3D]SamplerPoint
	Object3D      `collides:"true"`
}

func (p3d *Poly3D) GetPoints() []Point3D {
	return p3d.Points
}

func (p3d *Poly3D) GetSamplerPoints() map[Point3D]SamplerPoint {
	return p3d.SamplerPoints
}
