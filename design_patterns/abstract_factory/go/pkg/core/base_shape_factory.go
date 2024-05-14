package core

type BaseShapeFactory struct {
	Repo ShapeRepository
}

func (sf *BaseShapeFactory) MakeShape(shapeType ShapeType) (Shape, error) {
	shape, ok := sf.Repo.Shapes()[shapeType]
	if !ok {
		return nil, ErrShapeIsNotFound
	}
	return shape, nil
}

func (sf *BaseShapeFactory) GetShapeTypes() []ShapeType {
	shapeTypes := make([]ShapeType, 0, len(sf.Repo.Shapes()))
	for shapeType := range sf.Repo.Shapes() {
		shapeTypes = append(shapeTypes, shapeType)
	}
	return shapeTypes
}
