package app

type App struct{}

func (a *App) Start(shapeFactory ShapeFactory) {
	shape1, _ := shapeFactory.Make("circle")
	circle := shape1.(Shape)
	circle.Draw()
}
