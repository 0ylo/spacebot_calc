package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

//  const (
//     mixTax float64 = 0.0716
//     normalTax float64 = 0.08
// )

var dep, dayly, money, relax, first float64
var day int

//place for struct

func main() {
	var a int
	for {
		fmt.Printf("\n\n1 - Рассчитать процент \n2 - Коротко об условиях \n3 - Инструкции и подробнее о проекте\n4 - Задать вопрос \n0 - выход \n\n")
		fmt.Scanln(&a)
		switch {
		case a == 1:
			calculate()
		case a == 2:
			fmt.Println("\nИнвестиции под 8% в месяц с ежедневной выплатой процента (соответственно можно реинвестировать ежедневно!).\nТело депозита удерживается на 30 дней.\n")
		case a == 3:
			fmt.Println("\nПодробности и регистрация по ссылке https://teletype.in/@arousal/spacebot\n")
		case a == 4:
			fmt.Println("\nПо всем вопросам пишите на https://t.me/Surnev\n")
		case a == 0:
			fmt.Println("\nGoodbye!")
			os.Exit(0)
		default:
			fmt.Println("\nНет такого пункта меню\n")
		}
	}
}

func calculate() {
	//Задаем сумму депозита и срок вклада
	fmt.Println("\nПриветствую! Это калькулятор сложного процента для проекта SpaceBot!\n\nВведите сумму депозита в рублях:")
	fmt.Scanln(&dep)

	// Проверка положительной суммы
	for dep < 1 {
		fmt.Println("Введите пожалуйста значение больше нуля")
		fmt.Scanln(&dep)
	}
	money = dep

	// 3адаем количество месяцев
	fmt.Println("\nНа сколько месяцев вносим депозит?")
	days := calct(day)
	months := math.Round(float64(days) / 30)

	// Cчитаем доход за год без реинвестирования
	relax = dep + dep*0.0716*float64(months)
	if dep > 210000 {
		relax = dep + dep*0.08*float64(months)
	}
	fmt.Println("\nХорошо, без реинвестирования ваш депозит через", months, "месяцаев, составит:")
	fmt.Printf("%.2f\n", relax)

	// Cчитаем первую выплату по проценту
	first = money * 0.002327
	if money > 210000 {
		first = money * 0.0026
	}
	fmt.Printf("Ежедневно вам будет начисляться процент, начиная с\n%.2f\n", first)

	// Подсчет реинвестирования
	for i := 0; i < days; i++ {
		if money > 210000 {
			dayly = money * 0.002327
			money = dayly + money
		} else {
			dayly = money * 0.0026
			money = dayly + money
		}
	}
	fmt.Println("\n", "Ваш депозит через", months, "месяцяев при ежедневном реинвестировании:")
	fmt.Printf("%.2f\n", money)
}

func calct(day int) int {
	var now = time.Now()
	var month int
	fmt.Scanln(&month)
	for month < 1 {
		fmt.Println("Введите пожалуйста значение больше нуля")
		fmt.Scanln(&month)
	}
	var se = now.AddDate(0, month, 0)
	dur := se.Sub(now)
	allday := dur.Hours() / 24
	return int(allday)
}
