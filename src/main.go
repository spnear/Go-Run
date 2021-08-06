package main

import (
	"fmt"
	"math"
)

func area_rectangulo(base float64, altura float64) float64 {
	return base * altura
}

func area_trapecio(base_mayor float64, base_menor float64, altura float64) float64 {
	return (base_mayor + base_menor) / (2 * altura)
}

func area_circulo(radio float64) float64 {
	//A = πr 2
	const pi float64 = 3.1416
	return pi * math.Pow(radio, 2)
}

func double_return(a int, b int) (c, d int) {
	return a, a * b
}

func main() {

	fmt.Println("Hola GO :D")

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("Pi:", pi)
	fmt.Println("Pi2:", pi2)

	//Declaracion de varaibles
	base := 12          //Primera forma
	var altura int = 14 //Segunda forma
	var area int        //Go no compila si las variables no son usadas

	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	x := 10
	y := 20

	resultado_suma := x + y
	fmt.Println("Suma:", resultado_suma)

	resultado_resta := x - y
	fmt.Println("Resta:", resultado_resta)

	resultado_multiplicacion := x * y
	fmt.Println("Multiplicación:", resultado_multiplicacion)

	resultado_division := x / y
	fmt.Println("Suma:", resultado_division)

	resultado_modulo := x % y
	fmt.Println("Suma:", resultado_modulo)

	x++
	fmt.Println("Incremental", x)

	x--
	fmt.Println("Decremental", x)

	// Implementando área de rectángulo, trapecio y circulo
	// Rectangulo
	var base_rectangulo float64 = 10
	var altura_rectangulo float64 = 5
	fmt.Printf("Area del rectángulo: %.2f\n", area_rectangulo(base_rectangulo, altura_rectangulo))

	var base_menor_trapecio float64 = 5
	var base_mayor_trapecio float64 = 10
	var altura_trapecio float64 = 5
	fmt.Printf("Área de trapecio: %.2f\n", area_trapecio(base_mayor_trapecio, base_menor_trapecio, altura_trapecio))

	var radio_circulo float64 = 5
	fmt.Printf("Área circulo: %.2f\n", area_circulo(radio_circulo))

	value1, value2 := double_return(2, 3)
	//value1, _ := double_return(2, 3) //Podemos usar _ si no necesitamos ese valor (Recordar que go exige usar todas las variables declaradas)
	fmt.Printf("Valor 1 = %v. Valor 2 = %v\n", value1, value2)
}
