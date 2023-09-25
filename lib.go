package main

// Esta codigo es un proyecto, donde se debe crear una libreria.
// Este codigo trabaja con ARREGLOS no con SLICES.
// La diferencia radica en que los arreglos deben tener un tamano definido. Un slice es dinamico.
// El arreglo del main tiene una longitud de 10.
// Los elementos del arreglo pueden ser modificados en el main.
// El arreglo del main debe estar ORDENADO de MENOR a MAYOR.
// Este codigo no admite arreglos NO ORDENADOS.

import "fmt"

func agregarOrdenados(array [10]int, nuevoElemento int) {

	var arraySplit [10]int
	var arraySplit2 [10]int
	var resultado [11]int

	var longitudSplit = 0
	var contador = 0

	for i := 0; i <= len(array)-1; i++ {

		if array[i] <= nuevoElemento {
			arraySplit[longitudSplit] = array[i]
			longitudSplit++
		} else if array[i] > nuevoElemento {
			arraySplit2[i] = array[i]
		}

	}

	for j := 0; j < longitudSplit; j++ {
		resultado[j] = arraySplit[j]
		contador++
	}
	fmt.Println("Contador: ", contador)
	resultado[contador] = nuevoElemento

	if contador+1 != len(resultado) {
		for k := contador + 1; k <= len(resultado)-1; k++ {
			resultado[contador+1] = arraySplit2[contador]
			contador++
		}
	}
	fmt.Println("Arreglo original: ", array)
	fmt.Println("Arreglo modificado", resultado)
}

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	nuevoElemento := -5
	agregarOrdenados(array, nuevoElemento)
}
