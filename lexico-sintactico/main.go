package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LecturaTokens2() map[string]int {

	tokensBytes, err := os.ReadFile("tokens.txt")
	if err != nil {
		panic(err)
	}

	tokensString := string(tokensBytes)
	//fmt.Println(tokensString) PARA PRUEBAS
	tokensSplit := strings.Split(tokensString, "\n")
	//fmt.Println(tokensSplit) PARA PRUEBAS
	mapaTokens := make(map[string]int)

	for i := 0; i < len(tokensSplit); i++ {
		lineaActual := tokensSplit[i]

		partesDeLinea := strings.Split(lineaActual, ":")

		if len(partesDeLinea) == 2 {
			key := strings.TrimSpace(partesDeLinea[0])
			token := strings.TrimSpace(partesDeLinea[1])

			if token != "," {
				token = strings.Trim(token, `",`)
			}

			keyEntero, err := strconv.Atoi(key)
			if err == nil {
				mapaTokens[token] = keyEntero
			}
		}
	}
	return mapaTokens
}

func AnalizadorLexico2() []string {
	sqlBytes, err := os.ReadFile("sql.txt")
	if err != nil {
		panic(err)
	}

	sqlString := string(sqlBytes)
	fmt.Println("Lectura del archivo SQL:", sqlString)

	is := []string{}
	palabraActual := ""
	dentroDeComillas := false
	caracteresEspeciales := []byte{' ', ',', ';', '*', '+', '/', '=', '(', ')', '<', '>', '!', '%', '$', '@', '@', '&',
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
		} else if Contiene2(caracteresEspeciales, caracter) {
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

func Contiene2(slice []byte, elemento byte) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == elemento {
			return true
		}
	}
	return false
}

func LeerReglas() [][]int {
	var reglas [][]int

	contenido, err := os.ReadFile("sintactico.txt")
	if err != nil {
		panic(err)
	}

	lineas := strings.Split(string(contenido), "\n")
	for _, linea := range lineas {
		parts := strings.Split(linea, ":")
		if len(parts) == 2 {
			numero, err := strconv.Atoi(strings.TrimSpace(parts[0]))
			if err != nil {
				panic(err)
			}

			numerosStr := strings.Split(parts[1], ",")
			var numeros []int

			for _, numStr := range numerosStr {
				num, err := strconv.Atoi(strings.TrimSpace(numStr))
				if err != nil {
					panic(err)
				}
				numeros = append(numeros, num)
			}
			regla := append([]int{numero}, numeros...)
			reglas = append(reglas, regla)
		}
	}
	return reglas
}

func main() {
	mapaTokens := LecturaTokens2()
	datosSQL := AnalizadorLexico2()

	var sliceResultado []int
	var sliceCadenasUsuario []string

	bandera309 := false
	banderaNumerico := false

	for i := 0; i < len(datosSQL); i++ {
		buscar := datosSQL[i]

		valor, encontrado := mapaTokens[buscar]
		if buscar == "\n" {
			sliceResultado = append(sliceResultado, 30) //ASIGNAMOS EL TOKEN 30 PARA EL SALTO DE LINEA
		}
		if encontrado {
			sliceResultado = append(sliceResultado, valor)
			if valor == 309 {
				bandera309 = true
			}
			if valor == 13 || valor == 14 || valor == 10 {
				banderaNumerico = true
			}
		} else {
			buscar = strings.TrimSpace(buscar)
			if buscar != "" {
				if banderaNumerico {
					sliceResultado = append(sliceResultado, 997) //AGREGAMOS UN DATO NUMERICO
					sliceCadenasUsuario = append(sliceCadenasUsuario, buscar)
					banderaNumerico = false
				} else if bandera309 {
					sliceResultado = append(sliceResultado, 998) //AGREGAMOS UNA TABLA
					sliceCadenasUsuario = append(sliceCadenasUsuario, buscar)
					bandera309 = false
				} else {
					sliceResultado = append(sliceResultado, 999) //AGREGAMOS UNA CADENA DEL USUARIO
					sliceCadenasUsuario = append(sliceCadenasUsuario, buscar)
				}
			}
		}
	}

	reglas := LeerReglas()
	resultado := sliceResultado
	fmt.Println("Tokens del archivo SQL: ", resultado)
	fmt.Println("Arreglo de reglas: ", reglas)

	if ValidarSintaxis(resultado, reglas) {
		fmt.Println("Al buscar el arreglo de tokens dentro del arreglo de reglas, no se han encontrado errores.")
	} else {
		fmt.Println("No se encontre el arreglo de tokens dentro del arreglo de reglas")
	}
}

func ValidarSintaxis(resultadoSlice []int, reglas [][]int) bool {
	for i := 0; i < len(reglas); i++ {
		regla := reglas[i]
		if reflect.DeepEqual(resultadoSlice, regla) {
			return true
		}
	}
	return false
}
