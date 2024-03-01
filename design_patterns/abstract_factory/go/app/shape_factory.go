package app

type ShapeFactory interface {
	Make(shapeName string) (interface{}, error)
	GetShapeNames() []string
}
