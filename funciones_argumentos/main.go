package main

import (
	"fmt"
	"funciones_argumentos/greetings"
	"funciones_argumentos/ops"
)

func main() {
	greetings.HelloWorld()
	suma := ops.Sum(1, 2)
	fmt.Println(suma)
	resultado, residuo := ops.Div(5, 2)
	fmt.Println(resultado, residuo)
	sum, div := ops.SumAndDiv(7, 4)
	fmt.Println(sum, div)
	sub := ops.Substract(10, 2)
	fmt.Println(sub)
	sub1 := ops.Substract(10, nil)
	fmt.Println(sub1)
	sub2 := ops.Substract(10, "3")
	fmt.Println(sub2)
	fmt.Printf("Bienvenidas a %s version %d día %.1f", name, version, weekday)
}
