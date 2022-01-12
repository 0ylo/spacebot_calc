package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var dep, dayly, money, relax, first float64
	var day int

	//задаем сумму депозита и срок вклада
	fmt.Println("Приветствую! Это калькулятор сложного процента для проекта SpaceBot!")
	time.Sleep(1 * time.Second)
	fmt.Println("\n", "Введите сумму депозита в рублях:")
	fmt.Scanln(&dep)
	money = dep
	fmt.Println("На сколько месяцев вносим депозит?")
	days := calct(day)
	months := math.Round(float64(days) / 30)

	// считаем доход за год без реинвестирования
	if dep < 210000 {
		relax = dep + dep*0.0716*float64(months)
	} else {
		relax = dep + dep*0.08*float64(months)
	}
	fmt.Println("\n", "Хорошо, без реинвестирования ваш депозит через", months, "месяцаев, составит:")
	fmt.Printf("%.2f\n", relax)

	// считаем первую выплату по проценту
	if money < 210000 {
		first = money * 0.002327
	} else {
		first = money * 0.0026
	}
	fmt.Println("\n", "Ежедневно вам будет начисляться процент, начиная с")
	fmt.Printf("%.2f\n", first)

	// подсчет реинвестирования
	for i := 0; i < days; i++ {
		if money < 210000 {
			dayly = money * 0.002327
			money = dayly + money
		} else {
			dayly = money * 0.0026
			money = dayly + money
		}
	}
	fmt.Println("\n", "Ваш депозит через", months, "месяцяев при ежедневном реинвестировании:")
	fmt.Printf("%.2f\n", money)
	fmt.Println("\n")

	fmt.Println("Чтобы начать зарабатывать на инвестировании - регистрируйся по ссылке https://spaceltd.page.link/vsbKRBK6Mh66cFXbA \n", "По всем вопросам пишите на https://t.me/Surnev")
}

//подсчет календарных дней
func calct(day int) int {
	var now = time.Now()
	var month int
	fmt.Scanln(&month)
	var se = now.AddDate(0, month, 0)
	dur := se.Sub(now)
	allday := dur.Hours() / 24
	return int(allday)
}
