package forms

import (
	"errors"
	"fmt"
	"anythings/shape"
)

func CreateKvadrat(A float64) {
	fmt.Println("Введите значение стороны:")
	fmt.Scanf("%f\n", &A)
	PrintKvadrat(A)
}

//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ

func PrintKvadrat(A float64) {

	ploshad, err := calculatePloshadKvadrat(A)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Площадь квадрата составляет: %.2f см. кв.\n\n", ploshad)

	perimetr, err := calculatePerimetrKvadrat(A)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр квадрата составляет: %.2f см.\n", perimetr)
	shape.PringShape(forma, perimetr, ploshad)
}

func calculatePerimetrKvadrat(A float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("сторона не может быть меньше единицы")
	}

	return float64(A) * 4, nil
}

func calculatePloshadKvadrat(A float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("сторона не может быть меньше единицы")
	}

	return float64(A) * float64(A), nil
}

//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ//// КВАДРАТ