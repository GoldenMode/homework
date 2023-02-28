package forms

import (
	"errors"
	"fmt"
	"math"
	"anythings/shape"
)

func CreateTreugol(A, B, C float64) {
	fmt.Println("Введите значение первой стороны:")
	fmt.Scanf("%f\n", &A)
	fmt.Println("Введите значение второй стороны:")
	fmt.Scanf("%f\n", &B)
	fmt.Println("Введите значение третьей стороны:")
	fmt.Scanf("%f\n", &C)
	PrintTriugol(A, B, C)
}


///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК

func PrintTriugol(A, B, C float64) {

	perimetr, err := calculatePerimetrTriugol(A, B, C)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр треугольника составляет: %.2f см.\n", perimetr)

	polyPerimetr := calcPolyPerimetr(perimetr)

	visota := math.Sqrt(polyPerimetr*(polyPerimetr-A)*(polyPerimetr-B)*(polyPerimetr*C)) * 2 / A

	ploshad := calculatePloshadTriugol(A, visota)
	//fmt.Printf("Площадь треугольника составляет: %.2f см. кв.\n", ploshad)
	shape.PringShape(forma, perimetr, ploshad)
}

func calculatePerimetrTriugol(A, B, C float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("сторона треугольника не может быть меньше единицы")
	}
	if B <= 0 {
		return float64(0), errors.New("сторона треугольника не может быть меньше единицы")
	}
	if C <= 0 {
		return float64(0), errors.New("сторона треугольника не может быть меньше единицы")
	}
	return float64(A) + float64(B) + float64(C), nil
}

func calcPolyPerimetr(perimetr float64) float64 {
	return float64(perimetr) / 2
}

func calculatePloshadTriugol(A, visota float64) float64 {
	return (float64(A) * float64(visota)) / 2
}

///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК///// ТРЕУГОЛЬНИК
