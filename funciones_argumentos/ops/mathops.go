package ops

import "strconv"

//sum , miniscula privada para el paquete
//Sum , publica para los otros paquetes
func Sum(a, b int) int {
	return (a + b)
}

func Div(a, b int) (int, int) {
	return a / b, a % b
}

//esta funcion es para cuando tenemos pocos arreglos
func SumAndDiv(a, b int) (sum int, div int) {
	sum = Sum(a, b)
	div, _ = Div(a, b)
	return // naked return: retorno desnudo o retorno implicito
}

//el interface nos dice: puedo recibir cualquier tipo de dato incluso nulo
//seria como nuestro ultimo recurso
func Substract(a int, b interface{}) int {
	//if b == nil {
	//	return a
	//}
	var subB int
	total := 0 // si por alguna razon no entra al switch el valor retornado sera nil
	switch t := b.(type) {
	case int:
		total = a - t
	case int16:
		subB = int(t)
		total = a - subB
	case int32:
		subB = int(t)
		total = a - subB
	case int64:
		subB = int(t)
		total = a - subB
	case float32:
		subB = int(t)
		total = a - subB
	case string:
		subB, _ = strconv.Atoi(t) //ignoramos el error
		total = a - subB
	default:
		total = a
	}
	return total
}
