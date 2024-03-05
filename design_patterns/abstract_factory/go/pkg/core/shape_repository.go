package core

type ShapeRepository interface {
	Shapes() map[ShapeType]Shape
}
