package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Formato: calculadora <operacion> <num1> <num2>")
		os.Exit(1)
	}

	op := os.Args[1]
	n1, err1 := strconv.Atoi(os.Args[2])
	n2, err2 := strconv.Atoi(os.Args[3])

	/* validar numeros */
	if err1 != nil || err2 != nil {
		fmt.Println("Datos incorrectos, ingrese números enteros")
		return
	}

	switch op {
	case "suma":
		fmt.Printf("Resultado: %d\n", n1+n2)
	case "resta":
		fmt.Println("Resultado: ", n1-n2)
	case "multiplicacion":
		fmt.Println("Resultado: ", n1*n2)
	case "division":
		if n2 == 0 {
			fmt.Println("Error: divisón por cero")
			return
		} else {
			fmt.Printf("Resultado: %d\n", n1/n2)
		}
	default:
		fmt.Println("Operacion incorrecta")
	}
}
