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
	fmt.Printf("Bienvenidas a %s version %d día %.1f\n", name, version, weekday)

	//even or add
	numberParImpar := ops.ParImpar(9)
	fmt.Println(numberParImpar)

	//Ley de Ohm
	println("Ley Ohm 1")
	vol, _ := ops.OhmEasy(0, 2, 3)
	fmt.Printf("%.2f\n", vol)
	i, _ := ops.OhmEasy(6, 2, 0)
	fmt.Printf("%.2f\n", i)
	r, _ := ops.OhmEasy(6, 0, 3)
	fmt.Printf("%.2f\n", r)

	println("Ahora con Interface")
	v, _ := ops.Resolver(ops.Voltaje{R: 10, I: 2})
	fmt.Printf("%.2f", v)
}
