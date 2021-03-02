package main

import "fmt"

func main() {
	var dmg int = 40
	if dmg == 5 {
		fmt.Println("Царапина")
	}
	if dmg == 15 {
		fmt.Println("Лёгкое ранение")
	}
	if dmg == 40 {
		fmt.Println("Ранение,", "требуется бинт")
	}
	if dmg == 55 {
		fmt.Println("Тяжёлое ранение")
	}
	if dmg == 70 {
		fmt.Println("Кровотечение")
	} else if dmg == 5 {
		fmt.Println("5 dmg")
	} else if dmg == 15 {
		fmt.Println("15 dmg")
	} else if dmg == 40 {
		fmt.Println("40 dmg")
	} else if dmg == 55 {
		fmt.Println("55 dmg")
	} else if dmg == 70 {
		fmt.Println("70 dmg")
	}
}
