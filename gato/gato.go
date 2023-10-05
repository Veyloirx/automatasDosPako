package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func validarTablero(tablero string, jugadorDefault string) {
	// Verificar que el jugador predeterminado haya colocado su marcador en el tablero
	// Dependiendo del jugador predeterminado, se espera una X o una O en el tablero.
	if (jugadorDefault == "A" && !strings.Contains(tablero, "X")) || (jugadorDefault == "B" && !strings.Contains(tablero, "O")) {
		fmt.Printf("El jugador %s debe comenzar colocando su marcador en el tablero.\n", jugadorDefault)
		os.Exit(1)
	}
}

func actualizarLog(jugadorDefault string, tablero string, logContenido string) {
	movimientos := contarMovimientos(tablero)
	turno := determinarTurno(jugadorDefault, movimientos)
	movimientosA := obtenerMovimientos(tablero, "X")
	movimientosB := obtenerMovimientos(tablero, "O")

	// Construir el nuevo contenido del log
	nuevoLog := fmt.Sprintf("default = %s\nmovimientos = %d\nturno = %s\n", jugadorDefault, movimientos, turno)

	// Agregar movimientos_A al log si hay movimientos registrados
	if len(movimientosA) > 0 {
		nuevoLog += fmt.Sprintf("movimientos_A = [%s]\n", strings.Join(movimientosA, ", "))
	}

	// Agregar movimientos_B al log si hay movimientos registrados
	if len(movimientosB) > 0 {
		nuevoLog += fmt.Sprintf("movimientos_B = [%s]\n", strings.Join(movimientosB, ", "))
	}

	// Escribir el nuevo contenido en el archivo log.txt
	err := ioutil.WriteFile("log.txt", []byte(nuevoLog), 0644)
	if err != nil {
		fmt.Println("Error al escribir en log.txt:", err)
		os.Exit(1)
	}
}

func contarMovimientos(tablero string) int {
	movimientos := strings.Count(tablero, "X") + strings.Count(tablero, "O")
	return movimientos
}

func determinarTurno(jugadorDefault string, movimientos int) string {
	// Si el número total de movimientos es par, es el turno del jugador predeterminado, de lo contrario, es el turno del otro jugador.
	if movimientos%2 == 0 {
		return jugadorDefault
	}
	if jugadorDefault == "A" {
		return "B"
	}
	return "A"
}

func obtenerMovimientos(tablero string, marcador string) []string {
	// Buscar las coordenadas donde se encuentra el marcador en el tablero y almacenarlas en un slice.
	var movimientos []string
	tableroFilas := strings.Split(tablero, "\n")
	for i, fila := range tableroFilas {
		for j, c := range fila {
			if string(c) == marcador {
				movimientos = append(movimientos, fmt.Sprintf("[%d,%d]", i, j))
			}
		}
	}
	return movimientos
}

func mantenerLogExistente(logExistente string, nuevoLog string) string {
	// Mantener las líneas del log original que no han cambiado.
	lineasLogExistente := strings.Split(logExistente, "\n")
	lineasNuevoLog := strings.Split(nuevoLog, "\n")
	for i := 0; i < len(lineasLogExistente); i++ {
		lineaExistente := lineasLogExistente[i]
		lineaNueva := lineasNuevoLog[i]
		if strings.HasPrefix(lineaNueva, "movimientos_A =") || strings.HasPrefix(lineaNueva, "movimientos_B =") {
			// Mantener las líneas de movimientos sin cambios.
			lineasNuevoLog[i] = lineaExistente
		}
	}
	return strings.Join(lineasNuevoLog, "\n")
}

func obtenerJugadorDefault(logContenido string) string {
	// Buscar la línea "default =" en el contenido del log para obtener el jugador predeterminado.
	lineasLog := strings.Split(logContenido, "\n")
	for _, linea := range lineasLog {
		if strings.HasPrefix(linea, "default =") {
			parts := strings.Split(linea, "=")
			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}

func main() {
	// Leer el archivo de texto tablero.txt
	tableroBytes, err := os.ReadFile("tablero.txt")
	if err != nil {
		panic(err)
	}
	tablero := string(tableroBytes)

	// Leer el archivo de texto log.txt
	logBytes, err := os.ReadFile("log.txt")
	if err != nil {
		panic(err)
	}
	logContenido := string(logBytes)

	// Determinar el jugador predeterminado a partir del log existente
	jugadorDefault := obtenerJugadorDefault(logContenido)

	// Verificar que el tablero cumple con las condiciones
	validarTablero(tablero, jugadorDefault)

	// Actualizar el log con la información de los movimientos
	actualizarLog(jugadorDefault, tablero, logContenido)
}
