package main

import "fmt"

func sum(values ...int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total

}

func lista(values ...string) string {
	nombres := ""
	for _, nom := range values {
		nombres = nombres + nom + ","

	}
	return nombres
}

func getValues(x int) (int, int, int) {
	return 2 * x, 4 * x, 6 * x

}

func getValues2(x int) (double int, triple int, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5))

	fmt.Println(lista("Facundo", "Pablo", "Maria", "Carla"))

	fmt.Println(getValues(4))
	fmt.Println(getValues2(6))
}
