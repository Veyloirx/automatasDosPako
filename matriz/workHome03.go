package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//datosBytes representa la lectura en bytes
	datosBytes, err := os.ReadFile("datos.txt")
	if err != nil {
		panic(err)
	}
	datosString := string(datosBytes)
	datosSplit := strings.Split(datosString, "\n")

	var lineas int = 0
	for j := 0; j <= len(datosSplit)-1; j++ {
		lineas++
	}

	var longitud = len(datosSplit[0])
	for i := 1; i <= len(datosSplit)-1; i++ {
		if longitud != len(datosSplit[i]) {
			fmt.Println("Las longitudes son diferentes.")
			break
		}
	}

	var matriz [][]string
	
	for i := 0; i < lineas; i++ {
		fila := strings.Split(datosSplit[i], "")
		matriz = append(matriz, fila)
	}

	// Imprimir la matriz
	fmt.Println(matriz[0][1])
}
