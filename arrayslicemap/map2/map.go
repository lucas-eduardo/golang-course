package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":     11325.45,
		"Gabriel Silva": 15456.78,
		"Pefro Junior":  1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
