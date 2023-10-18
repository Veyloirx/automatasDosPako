package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LecturaTokens() map[string]int {
	tokensBytes, err := os.ReadFile("tokens.txt")
	if err != nil {
		panic(err)
	}
	tokensString := string(tokensBytes)
	tokensSplit := strings.Split(tokensString, "\n")

	mapaTokens := make(map[string]int)

	for i := 0; i < len(tokensSplit); i++ {
		lineaActual := tokensSplit[i]

		partesDeLinea := strings.Split(lineaActual, ":")
		if len(partesDeLinea) == 2 {
			key := strings.TrimSpace(partesDeLinea[0])
			token := strings.TrimSpace(partesDeLinea[1])

			token = strings.Trim(token, `",`)

			keyEntero, err := strconv.Atoi(key)
			if err == nil {
				mapaTokens[token] = keyEntero
			}
		}
	}
	return mapaTokens
}

func AnalizadorLexico() []string {
	sqlBytes, err := os.ReadFile("sql.txt")
	if err != nil {
		panic(err)
	}

	is := []string{}
	palabraActual := ""
	dentroDeComillas := false
	caracteresEspeciales := []byte{'.', '_', ',', ' ', ';', '*', '+', '/', '=', '(', ')', '<', '>', '!', '%', '$', '@', '@', '&',
		'|', '^', '`', '~', '?', ':', '"', '\'', '[', ']', '\n'}

	for i := 0; i < len(sqlBytes); i++ {
		caracter := sqlBytes[i]

		if caracter == '\'' {
			if !dentroDeComillas {
				is = append(is, "'")
				dentroDeComillas = true
			} else {
				is = append(is, palabraActual)
				is = append(is, "'")
				dentroDeComillas = false
				palabraActual = ""
			}
		} else if dentroDeComillas {
			if caracter == '\\' {
				i++
				continue
			}
			palabraActual += string(caracter)
		} else if Contiene(caracteresEspeciales, caracter) {
			if palabraActual != "" {
				is = append(is, palabraActual)
				palabraActual = ""
			}
			is = append(is, string(caracter))
		} else {
			palabraActual += string(caracter)
		}
	}

	if palabraActual != "" {
		is = append(is, palabraActual)
	}

	//PARA PRUEBAS
	//for _, token := range is {
	//	if token != " " {
	//		fmt.Printf("%q, ", token)
	//	}
	//}
	return is
}

func Contiene(slice []byte, elemento byte) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == elemento {
			return true
		}
	}
	return false
}

func MarcarPalabrasUsuario(sliceResultado []int) []int {
	resultado := make([]int, 0)
	bandera := false

	for i := 0; i < len(sliceResultado); i++ {
		token := sliceResultado[i]

		if token == 27 {
			bandera = !bandera
			if bandera {
				resultado = append(resultado, 999)
			}
		} else if bandera {
			resultado = append(resultado, 999)
		} else {
			resultado = append(resultado, token)
		}
	}

	return resultado
}

func main() {

	//analizadorLexico() PARA PRUEBAS
	mapaTokens := LecturaTokens()
	datosSQL := AnalizadorLexico()

	var sliceResultado []int
	for i := 0; i < len(datosSQL); i++ {
		evaluar := datosSQL[i]

		valor, encontrado := mapaTokens[evaluar]
		if encontrado {
			sliceResultado = append(sliceResultado, valor)
		}
	}
	resultado := MarcarPalabrasUsuario(sliceResultado)
	fmt.Println(resultado)
	//fmt.Println(sliceResultado) slice sin palabas del usuario.
}
