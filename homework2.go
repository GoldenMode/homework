package main

import (
	"fmt"
	"math"
)

func main() {
	var n int = 5

	createMassiv(n)

	var todoList = []string{
		"bread",
		"milk",
		"butter",
		"cheese",
		"bottle of water",
		"eggs",
	}

	fmt.Println("")

	tasks := todoList[0:6]
	showTasksBefore(tasks)

	fmt.Println("\n ----- ПОСЛЕ changeTasks() ----\n")

	changeTasks(tasks)
	showTasksAfter(tasks)

}

func changeTasks(tasks []string) {

	for w := 1; w <= len(tasks)-1; w++ {
		for i := 0; i < len(tasks)-w; i++ {
			q := tasks[i]
			tasks[i] = tasks[i+1]
			tasks[i+1] = q

		}
	}
	return
}

func showTasksAfter(tasks []string) {
	for i := range tasks {
		fmt.Println(tasks[i])
	}

	fmt.Println("")
}
func showTasksBefore(tasks []string) {
	for i := range tasks {
		fmt.Println(tasks[i])
	}

	fmt.Println("")
}

func createMassiv(n int) {
	for i := 1; i <= n; i++ {
		c := math.Pow(2, float64(i))
		fmt.Printf("%.1f\n", c)
	}
	fmt.Println("")
	return
}
