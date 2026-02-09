package main

import "fmt"

func main() {

	// En go las variables se escriben con CAMELCASE
	/*
		Los comentarios de varias lineas
		se representan con * y /
	*/

	//Variables

	//inferencia
	//edad := 25

	//INTEGERS

	//var edad int //variable sin valor
	var edad uint //unsigned integer, varibale sin signo negativo

	/*
		unit 8,16,32,64 go esta pensando para optimizar, se evalua cuantos bit se necesitan.
	*/
	//edad =-25 // go genera solo el error

	edad = 25
	fmt.Println(edad)
	edad = 123
	fmt.Println(edad)
	fmt.Println(CurrentYear - edad)

	//FLOATING
	var myFloat float32 //precision simple -> 7 dígitos decimales
	myFloat = 25.4
	fmt.Println(myFloat)
	var myFloat2 float64 // presicion doble -> 15 digitos decimales
	myFloat2 = 25.4
	fmt.Println(myFloat2)

	//STRING
	nombre := "Juliana"
	nombre = nombre + " Dev"
	fmt.Println(nombre)

	//BOOLEAN
	esMayorQue := true //camelcase NO la quiero prestar a ningun otro paquete, variable privada
	//EsMayorQue := true //pascalcase la quiero prestar a otros paquetes variable publica

	fmt.Println(esMayorQue)

}
