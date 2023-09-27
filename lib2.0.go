package main

// Este codigo acepta cualquier tipo de arreglo que se le envie
// Inserta un valor en la posicion anterior al numero mayor a este

import "fmt"

func agregar(array [10]int, nuevoElemento int) {

	var arraySplit [10]int
	var resultado [11]int

	var longitudSplit = 0
	var contador = 0

	for i := 0; i <= len(array); i++ {

		if array[i] > nuevoElemento {
			break
		}
		arraySplit[longitudSplit] = array[i]
		longitudSplit++
	}
	fmt.Println(arraySplit)

	for j := 0; j < longitudSplit; j++ {
		resultado[j] = arraySplit[j]
		contador++
	}
	resultado[contador] = nuevoElemento

	if contador+1 != len(resultado) {
		for k := contador + 1; k <= len(resultado)-1; k++ {
			resultado[contador+1] = array[contador]
			contador++
		}
	}
	fmt.Println("Arreglo original: ", array)
	fmt.Println("Arreglo modificado", resultado)
}

func main() {
	array := [10]int{1, 9, 6, 2, 12, 76, 117, 120, 2, 190}

	nuevoElemento := -1
	agregar(array, nuevoElemento)
}
