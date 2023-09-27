package main

//generar un diccionario de datos tomando como datos tomando como keys = arreglo y como data el arreglo data
//data debe ser igual o mayor a arreglo

func main() {
	var arreglo = [5]int{0, 1, 2, 3, 4}
	var data = [5]string{"saludo", "sour", "interger", "plutonium", "quack"}

	var diccionario = map[string]int{}
	for i := 0; i <= len(arreglo)-1; i++ {
		if len(data) > len(arreglo) {
			break
		}
		diccionario[data[i]] = arreglo[i]
	}

}
