package forms

import (
	
	"fmt"
)

var forma string
var A, B, C float64

func CreateForma() {
	fmt.Println("Для подсчёта фигуры по Вашим данным выберите её форму [Круг, Квадрат, Прямоугольник, Треугольник]")
	fmt.Scanf("%s\n", &forma)
	TakeForma(forma)
}

func TakeForma(forma string) {
	switch forma {
	case "Круг", "круг", "круГ":
		fmt.Printf("Выбранной фигурой является %s, начинаем подсчёт\n", forma)
		CreateKrug(A)
	case "Квадрат", "квадрат", "квадраТ":
		fmt.Printf("Выбранной фигурой является %s, начинаем подсчёт\n", forma)
		CreateKvadrat(A)
	case "Прямоугольник", "прямоугольник", "прямоугольниК":
		fmt.Printf("Выбранной фигурой является %s, начинаем подсчёт\n", forma)
		CreatePryamo(A, B)
	case "Треугольник", "треугольник", "треугольниК":
		fmt.Printf("Выбранной фигурой является %s, начинаем подсчёт\n", forma)
		CreateTreugol(A, B, C)
	default:
		fmt.Println("Вы выбрали не существующую фигуру или допустили ошибку. Попробуйте еще раз!")
	}
}