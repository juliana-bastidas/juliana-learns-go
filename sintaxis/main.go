package main

import "fmt"

func main() {

	// En go las variables se escriben con CAMELCASE, myVariable
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

	// Math operators
	//arithmetic operators

	fmt.Println(1 + 2) //suma
	fmt.Println(3 - 2) //resta
	fmt.Println(3 * 2) //multiplicacion
	fmt.Println(4 / 2) // division =2
	fmt.Println(4 % 2) // modulo = 0
	//comparison operators
	fmt.Println(1 == 2) //igualdad
	fmt.Println(1 != 2) //desigualdad
	fmt.Println(1 >= 2) //mayor igual que
	fmt.Println(1 <= 2) //menor igual que

	//logical operators , el sombreado naranja es una advertencia no un error por lo tanto deja compilar
	fmt.Println(true && true) //AND
	fmt.Println(true || true) // OR
	fmt.Println(!true)        //NOT

	/*
		FLOW CONTROL
	*/

	//if

	if edad < 18 {
		fmt.Println("No puedes entrar, eres menor de edad")
	} else {
		fmt.Println("Adelante")
	}

	//for
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}

	//while: NO EXISTE EN GO, el for se comporta como un while
	total := 0
	for total < 10 {
		fmt.Println(total)
		total++
	}

	//for ...range
	nombres := []string{"a", "b", "c"}
	for _, item := range nombres {
		fmt.Println(item)
	}

	//swith

	mes := "Agosto"
	switch mes {
	case "Enero":
		fmt.Println("Feliz año nuevo")
	case "Diciembre":
		fmt.Println("Fin de año")
	default:
		fmt.Println("Cualquier otro mes")
	}
	dia := `asd {edad \n asd}`
	fmt.Println(dia)

	fmt.Printf("Tengo %d años de edad \n", edad)
	formattedDate := fmt.Sprintf("Tengo %d años de edad", edad)
	fmt.Println(formattedDate)
	/*
		fmt.Printf: envia inmediatamente el string con formato a la consola
		espera el formato completo incluso los saltos de linea.
		fmt.Sprintf: asigna el resultado a una variable, para despues utilizar esa variable
	*/
}
