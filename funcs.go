package main

import "fmt"

func main() {
	const ft float64 = 0.3048 // это миллиметры

	var m = 66.12 // это футы
	{
		fmt.Printf("%.3f", m*ft) // это вычисление кол-ва миллиметров в указанных футах
		fmt.Println(" m")
		//fmt.Println(" m") - для удобства)
	}
}
