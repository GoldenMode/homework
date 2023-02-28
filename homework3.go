package main

import (
	"errors"
	"fmt"
	"math"
)

type employee struct {
	name     string
	perimetr float64
	ploshad  float64
}

func newEmployee(name string, perimetr, ploshad float64) employee {
	return employee{
		name:     name,
		perimetr: perimetr,
		ploshad:  ploshad,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Название фигуры: %s\nПериметр: %f\nПлощадь: %f\n", e.name, e.perimetr, e.ploshad)
}

func main() {

	printPryamo(5, 4)
	printTriugol(5, 6, 7)

}

func printPryamo(dlina, shirina float64) {
	//fmt.Printf("\nДлина прямоугольника = %.2f см.\n", dlina)
	//fmt.Printf("Ширина прямоугольника = %.2f см.\n", shirina)
	//fmt.Println("Формула для расчёта площади прямоугольника S=a*b \n")

	ploshadPryamo, err := calculatePloshadPryamo(dlina, shirina)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Площадь прямоульника составляет: %.2f см. кв.\n\n", ploshadPryamo)
	//fmt.Println("Формула для расчёта периметра прямоугольника P=2*(a+b) \n")

	perimetrPryamo, err := calculatePerimetrPryamo(dlina, shirina)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр прямоугольника составляет: %.2f см.\n", perimetrPryamo)

	employee1 := struct {
		name     string
		dlina    float64
		shirina  float64
		perimetr float64
		ploshad  float64
	}{
		name:     "Прямоугольник",
		dlina:    dlina,
		shirina:  shirina,
		perimetr: perimetrPryamo,
		ploshad:  ploshadPryamo,
	}
	fmt.Printf("%+v\n", employee1)

	employee12 := newEmployee(employee1.name, perimetrPryamo, ploshadPryamo)

	fmt.Printf("%s\n", employee12.getInfo())
}

func calculatePerimetrPryamo(dlina, shirina float64) (float64, error) {
	if dlina <= 0 {
		return float64(0), errors.New("Длина не может быть меньше единицы!")
	}
	if shirina <= 0 {
		return float64(0), errors.New("Ширина не может быть меньше единицы!")
	}

	return (float64(dlina) + float64(shirina)) * 2, nil
}

func calculatePloshadPryamo(dlina, shirina float64) (float64, error) {
	if dlina <= 0 {
		return float64(0), errors.New("Длина не может быть меньше единицы!")
	}
	if shirina <= 0 {
		return float64(0), errors.New("Ширина не может быть меньше единицы!")
	}
	return float64(dlina) * float64(shirina), nil
}

func printTriugol(A, B, C float64) {
	//fmt.Printf("\nА сторона треугольника = %.2f см.\n", A)
	//fmt.Printf("В сторона треугольника = %.2f см.\n", B)
	//fmt.Printf("С сторона треугольника = %.2f см.\n\n", C)
	//fmt.Println("Формула для нахождения периметра треугольника P=A+B+C")

	perimetrTriugol, err := calculatePerimetrTriugol(A, B, C)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//fmt.Printf("Периметр треугольника составляет: %.2f см.\n", perimetrTriugol)

	//fmt.Println("Формула для нахождения полупериметра треугольника p=(A+B+C)/2")

	polyPerimetr := calcPolyPerimetr(perimetrTriugol)
	//fmt.Printf("Полупериметр треугольника составляет: %.2f см., он понадобится для нахождения площади\n", polyPerimetr)

	//fmt.Println("Формула для нахождения высоты треугольника на сторону А h=(2*sqrt(p*(p-A)*(p-B)*(p-C)))/2")

	visota := math.Sqrt(polyPerimetr*(polyPerimetr-A)*(polyPerimetr-B)*(polyPerimetr*C)) * 2 / A
	//fmt.Printf("Высота направленная под прямым углом на сторону A: %.2f см.\n", visota)

	//fmt.Println("Формула для нахождения площади треугольника S=A*h/2")

	ploshadTriugol := calculatePloshadTriugol(A, visota)
	//fmt.Printf("Площадь треугольника составляет: %.2f см. кв.\n", ploshadTriugol)

	employee2 := struct {
		name     string
		storona1 float64
		storona2 float64
		storona3 float64
		perimetr float64
		ploshad  float64
	}{
		name:     "Треугольник",
		storona1: A,
		storona2: B,
		storona3: C,
		perimetr: perimetrTriugol,
		ploshad:  ploshadTriugol,
	}
	fmt.Printf("%+v\n", employee2)

	employee123 := newEmployee(employee2.name, perimetrTriugol, ploshadTriugol)

	fmt.Printf("%s\n", employee123.getInfo())
}

func calculatePerimetrTriugol(A, B, C float64) (float64, error) {
	if A <= 0 {
		return float64(0), errors.New("Сторона треугольника не может быть меньше единицы!")
	}
	if B <= 0 {
		return float64(0), errors.New("Сторона треугольника не может быть меньше единицы!")
	}
	if C <= 0 {
		return float64(0), errors.New("Сторона треугольника не может быть меньше единицы!")
	}
	return float64(A) + float64(B) + float64(C), nil
}

func calcPolyPerimetr(perimetrTriugol float64) float64 {
	return float64(perimetrTriugol) / 2
}

func calculatePloshadTriugol(A, visota float64) float64 {
	return (float64(A) * float64(visota)) / 2
}
