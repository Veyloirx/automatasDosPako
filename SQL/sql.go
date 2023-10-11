package main

import (
	"fmt"
	"os"
)

func analizadorLexico(expresion string) []string {
	is := []string{}
	palabraActual := ""

	for i := 0; i < len(expresion); i++ {
		caracter := expresion[i]

		if caracter == ' ' || caracter == '\n' || caracter == ',' {
			if palabraActual != "" {
				is = append(is, palabraActual)
			}

			if caracter == '\n' {
				is = append(is, "\n")
			} else if caracter == ',' {
				is = append(is, ",")
			}

			palabraActual = ""
		} else if caracter == '=' || caracter == '\'' {
			if palabraActual != "" {
				is = append(is, palabraActual)
			}

			is = append(is, string(caracter))
			palabraActual = ""
		} else {
			palabraActual += string(caracter)
		}
	}

	if palabraActual != "" {
		is = append(is, palabraActual)
	}

	return is
}

func main() {
	datosBytes, err := os.ReadFile("sql.txt")
	if err != nil {
		panic(err)
	}
	datos := string(datosBytes)

	is := analizadorLexico(datos)

	for _, token := range is {
		fmt.Printf("%q, ", token)
	}
}
