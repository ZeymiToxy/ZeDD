package main

import "fmt"

func main() {
	var dmg int = 15
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
		fmt.Println("Получено 5 урона")
	} else if dmg == 15 {
		fmt.Println("Получено 15 урона")
	} else if dmg == 40 {
		fmt.Println("Получено 40 урона")
	} else if dmg == 55 {
		fmt.Println("Получено 55 урона")
	} else if dmg == 70 {
		fmt.Println("Получено 70 урона")
	}
}
