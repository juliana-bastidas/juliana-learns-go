package ops

import (
	"fmt"
	"strconv"
)

// sum , miniscula privada para el paquete
// Sum , publica para los otros paquetes
func Sum(a, b int) int {
	return (a + b)
}

func Div(a, b int) (int, int) {
	return a / b, a % b
}

// esta funcion es para cuando tenemos pocos arreglos
func SumAndDiv(a, b int) (sum int, div int) {
	sum = Sum(a, b)
	div, _ = Div(a, b)
	return // naked return: retorno desnudo o retorno implicito
}

// el interface nos dice: puedo recibir cualquier tipo de dato incluso nulo
// seria como nuestro ultimo recurso
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

func ParImpar(a int) string {
	if a%2 == 0 {
		return "Es Par"
	} else {
		return "Es impar"
	}
}

func OhmEasy(v, i, r float64) (float64, error) {
	if v == 0 {
		return (i) * (r), nil
	}
	if i == 0 {
		if r == 0 {
			return 0, fmt.Errorf("división por cero")
		}
		return v / r, nil
	}
	if r == 0 {
		if i == 0 {
			return 0, fmt.Errorf("división por cero")
		}
		return v / i, nil
	}
	return 0, fmt.Errorf("debe faltar exactamente uno")
}

// Funcion con interface

type ohmCalculator interface {
	Calcular() (float64, error) //cualquier tipo que tenga calcular sirve
}
type Voltaje struct {
	R float64
	I float64
}

func (v Voltaje) Calcular() (float64, error) { //Voltaje tiene una función llamada Calcular
	return v.R * v.I, nil //V = R * I
}

type Corriente struct {
	V float64
	R float64
}

func (c Corriente) Calcular() (float64, error) {
	if c.R == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return c.V / c.R, nil
}

type Resistencia struct {
	V float64
	I float64
}

func (r Resistencia) Calcular() (float64, error) {
	if r.I == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return r.V / r.I, nil
}

func Resolver(o ohmCalculator) (float64, error) {
	return o.Calcular()
}
