package forms

import (
	"errors"
	"fmt"
	"anythings/shape"
)

func CreatePryamo(A, B float64) {
	fmt.Println("Введите значение для длины:")
	fmt.Scanf("%f\n", &A)
	fmt.Println("Введите значение для ширины:")
	fmt.Scanf("%f\n", &B)
	PrintPryamo(A, B)
}

///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК

func PrintPryamo(A, B float64) {

	ploshad, err := calculatePloshadPryamo(A, B)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Площадь прямоульника составляет: %.2f см. кв.\n\n", ploshad)

	perimetr, err := calculatePerimetrPryamo(A, B)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр прямоугольника составляет: %.2f см.\n", perimetr)
	shape.PringShape(forma, perimetr, ploshad)
}

func calculatePerimetrPryamo(A, B float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("длина не может быть меньше единицы")
	}
	if B <= 0 {
		return float64(0), errors.New("ширина не может быть меньше единицы")
	}

	return (float64(A) + float64(B)) * 2, nil
}

func calculatePloshadPryamo(A, B float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("длина не может быть меньше единицы")
	}
	if B <= 0 {
		return float64(0), errors.New("ширина не может быть меньше единицы")
	}
	return float64(A) * float64(B), nil
}

///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК///// ПРЯМОУГОЛЬНИК