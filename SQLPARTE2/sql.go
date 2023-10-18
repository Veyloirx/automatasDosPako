package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func LecturaTokens() map[int]string {
	tokensBytes, err := os.ReadFile("tokens.txt")
	if err != nil {
		panic(err)
	}
	tokensString := string(tokensBytes)
	tokensSplit := strings.Split(tokensString, "\n")

	mapaTokens := make(map[int]string)

	for i := 0; i < len(tokensSplit); i++ {
		lineaActual := tokensSplit[i]

		partesDeLinea := strings.Split(lineaActual, ":")
		if len(partesDeLinea) == 2 {
			key := strings.TrimSpace(partesDeLinea[0])
			token := strings.TrimSpace(partesDeLinea[1])

			token = strings.ReplaceAll(token, `"`, "")

			keyEntero, err := strconv.Atoi(key)
			if err == nil {
				mapaTokens[keyEntero] = token
			}
		}
	}
	return mapaTokens
}

func separadorLexico() []string {

	expresion, err := os.ReadFile("sql.txt")
	if err != nil {
		panic(err)
	}

	slice := []string{} //slice vacio
	palabraActual := ""

	for i := 0; i < len(expresion); i++ {
		caracter := expresion[i]

		if caracter == ' ' || caracter == '\n' || caracter == ',' {
			if palabraActual != "" {
				slice = append(slice, palabraActual)
			}
			if caracter == '\n' {
				slice = append(slice, "\n")
			} else if caracter == ',' {
				slice = append(slice, ",")
			}
			palabraActual = ""
		} else if caracter == '=' || caracter == '\'' || caracter == '(' || caracter == ')' || caracter == '*' {
			if palabraActual != "" {
				slice = append(slice, palabraActual)
			}
			slice = append(slice, string(caracter))
			palabraActual = ""
		} else {
			palabraActual += string(caracter)
		}
	}

	if palabraActual != "" {
		slice = append(slice, palabraActual)
	}

	return slice
}

func imprimirTokens(mapa map[int]string, tokens []string) {
	for _, token := range tokens {
		limpio := limpiarToken(token)
		valor, encontrado := buscarValorEnMapa(mapa, limpio)
		if encontrado {
			fmt.Printf("%s: %d\n", token, valor)
		} else {
			fmt.Printf("Token no encontrado en el mapa: %s\n", token)
		}
	}
}

func buscarValorEnMapa(mapa map[int]string, token string) (int, bool) {
	limpio := limpiarToken(token)
	for key, value := range mapa {
		limpioValor := limpiarToken(value)
		if limpioValor == limpio {
			return key, true
		}
	}
	return 0, false
}

func limpiarToken(token string) string {
	var limpio strings.Builder
	for _, r := range token {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			limpio.WriteRune(unicode.ToUpper(r))
		}
	}
	return limpio.String()
}

func main() {
	mapaTokens := LecturaTokens()
	tokens := separadorLexico()
	imprimirTokens(mapaTokens, tokens)
}
