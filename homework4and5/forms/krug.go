package forms

import (
	"errors"
	"fmt"
	"anythings/shape"
)

const pi = 3.1415

func CreateKrug(A float64) {
	fmt.Println("Введите значение радиуса Вашего круга:")
	fmt.Scanf("%f\n", &A)
	PrintKrug(A)
	}

//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ

func PrintKrug(A float64) {

	ploshad, err := calculatePloshadKrug(A)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Площадь квадрата составляет: %.2f см. кв.\n\n", ploshad)

	perimetr, err := calculatePerimetrKrug(A)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр квадрата составляет: %.2f см.\n", perimetr)
	shape.PringShape(forma, perimetr, ploshad)
}

func calculatePerimetrKrug(A float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("радиус не может быть меньше единицы")
	}

	return float64(A) * 2 * pi, nil
}

func calculatePloshadKrug(A float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("радиус не может быть меньше единицы")
	}

	return float64(A) * pi * float64(A), nil
}

//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ//// КРУГ