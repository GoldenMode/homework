package shape

import "fmt"

type shape struct {
	name     string
	perimetr float64
	ploshad  float64
}

func newShape(forma string, perimetr, ploshad float64) shape {
	return shape{
		name:     forma,
		perimetr: perimetr,
		ploshad:  ploshad,
	}
}

func (s shape) getInfo() string {
	return fmt.Sprintf("Название фигуры: %s\nПериметр: %0.2f см.\nПлощадь: %0.2f см. кв.\n", s.name, s.perimetr, s.ploshad)
}

func PringShape(forma string, perimetr, ploshad float64) {
	shapeFigure := newShape(forma, perimetr, ploshad)

	fmt.Printf("%s\n", shapeFigure.getInfo())
}
