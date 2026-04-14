package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	//args := os.Args
	if len(os.Args) < 2 {
		log.Fatal("No estas pasando los argumentos necesarios") // genera un print y exit del programa
	} else {
		switch action := os.Args[1]; action {
		case "hello":
			helloWorld()

		//go run main.go suma 1 2
		case "suma":
			//a, _ := strconv.Atoi(os.Args[2])
			//b, _ := strconv.Atoi(os.Args[3])

			var a int
			var b int

			fmt.Println("Introduce el primer valor:")
			fmt.Scanf("%d\n", &a)
			fmt.Println("Introduce el segundo valor:")
			fmt.Scanf("%d\n", &b)
			fmt.Printf("El resultado de la suma es: %d \n", sum(a, b))
		}
	}
	PrintManyArgs(10, "User logs in", "user id: qwerty12345")

}

// para ejecutarlo sin go.mod
// % go fun main.go hello
func helloWorld() {
	fmt.Println("hello desde Greeting")
}

func sum(a, b int) int {
	return (a + b)
}

func PrintManyArgs(level int, message string, details ...interface{}) {
	fmt.Printf("%d:::%s<%v>\n", level, message, details)
}
