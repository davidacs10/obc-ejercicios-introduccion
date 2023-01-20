package main

import "fmt"

func main() {
	//if
	numeroIf := -100

	if numeroIf > 0 {
		fmt.Println("Es positivo")
	}
	if numeroIf < 0 {
		fmt.Println("Es negativo")
	}
	if numeroIf == 0 {
		fmt.Println("Es cero")
	}

	//bucle while
	numeroWhile := -3
	// for {
	// 	if numeroWhile < 3 {
	// 		numeroWhile++
	// 		fmt.Println(numeroWhile)
	// 	}
	// }

	//do while
	for numeroWhile < 3 {
		numeroWhile++
		if numeroWhile != 3 {
			fmt.Println(numeroWhile)
		}
		break
	}

	//Bucle For
	for numeroFor := 0; numeroFor <= 3; numeroFor++ {
		fmt.Println(numeroFor)
	}

	//Switch
	var estacion = "otono"

	switch {
	case estacion == "primavera" || estacion == "Primavera":
		fmt.Println("Esta estacion es: ", estacion)
	case estacion == "verano" || estacion == "Verano":
		fmt.Println("Esta estacion es: ", estacion)
	case estacion == "otoño" || estacion == "Otoño":
		fmt.Println("Esta estacion es: ", estacion)
	case estacion == "invierno" || estacion == "Invierno":
		fmt.Println("Esta estacion es: ", estacion)
	default:
		fmt.Println("Esto no es una estacion")
	}
}
