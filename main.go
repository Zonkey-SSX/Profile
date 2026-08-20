package main

import "fmt"

func main() {

	fmt.Println("Привет альпинист получи свою карточку для покорения вершин  :) ")
	fmt.Println("Для начала расскажи мне о себе")

	name, age, weight, height, expirience, helmet, boots, insurance, backpackWeight := InputUser()

	ageOk := checkAge(age)

	expirienceOk := checkExpirience(expirience)

	equipmentOk := checkEquipment(helmet, boots, insurance)

	mountName, mountHeight, temperature := inputMountInfo()

	empty, difficultyLevel := mountainDifficultyLevel(ageOk, equipmentOk, expirienceOk, mountHeight, temperature)

	alpCard(name, age, weight, height, expirience, helmet, boots, insurance, empty, backpackWeight, difficultyLevel, mountName, mountHeight, temperature)

	/*result, weight := inputProfile()
	if !result {
		return
	}

	resultequip, backpackWeight := equipmentChek()
	if !resultequip {
		return
	}

	if !calculateWeight(weight, backpackWeight) {
		return
	}

	weatherCheck() */
}

func InputUser() (string, int, float64, int, int, string, string, string, float64) {

	var name string
	var age int
	var weight float64
	var height int
	var expirience int
	var helmet string
	var boots string
	var insurance string
	var backpackWeight float64

	fmt.Println("Как тебя зовут?")
	fmt.Scan(&name)

	fmt.Println("Сколько тебе лет?")
	fmt.Scan(&age)

	if !checkAge(age) {
		return "", 0, 0, 0, 0, "", "", "", 0
	}

	fmt.Println("Какой у тебя вес?")
	fmt.Scan(&weight)

	fmt.Println("Какой у тебя рост?")
	fmt.Scan(&height)

	fmt.Println("Какой у тебя опыт покорений гор в годах?")
	fmt.Scan(&expirience)

	if !checkExpirience(expirience) {
		return "", 0, 0, 0, 0, "", "", "", 0
	}

	fmt.Println("Есть ли у тебя каска? (да/нет)")
	fmt.Scan(&helmet)

	fmt.Println("Есть ли у тебя горные ботинки? (да/нет)")
	fmt.Scan(&boots)

	fmt.Println("Есть ли у тебя страховка? (да/нет)")
	fmt.Scan(&insurance)

	if !checkEquipment(helmet, boots, insurance) {
		return "", 0, 0, 0, 0, "", "", "", 0
	}

	fmt.Println("какой вес у твоего рюказака в кг?")
	fmt.Scan(&backpackWeight)

	if !checkWeight(weight, backpackWeight) {
		return "", 0, 0, 0, 0, "", "", "", 0
	}

	return name, age, weight, height, expirience, helmet, boots, insurance, backpackWeight
}

func checkAge(age int) bool {

	if age >= 18 {
	} else {
		fmt.Println("Возвращайся когда стукнет 18 ")
		return false
	}

	return true

}

func checkExpirience(expirience int) bool {

	if expirience > 1 {
		fmt.Println("Новые верха ждут :)")
	} else {
		fmt.Println("Наберись опыта")
		return false
	}
	return true
}

func checkEquipment(helmet string, boots string, insurance string) bool {

	switch {
	case helmet == "да":
	case helmet == "нет":
		fmt.Println("Тебе нужна каска")
		return false

	default:

		fmt.Println("Нужно вводить только да или нет")
		return false
	}

	switch {
	case boots == "да":
	case boots == "нет":
		fmt.Println("Тебе нужны горные ботинки")
		return false

	default:
		fmt.Println("Нужно вводить только да или нет")
		return false
	}

	switch {
	case insurance == "да":
	case insurance == "нет":
		fmt.Println("Тебе нужна страховка")
		return false
	default:

		fmt.Println("Используй только да/нет")
		return false
	}

	return true
}

func checkWeight(weight float64, backpack float64) bool {

	var sum float64

	sum = weight * 0.25

	switch {
	case backpack > sum:
		fmt.Println("С такой тяжестью не взлетишь")
		return false

	default:
	}

	return true
}

func inputMountInfo() (string, int, int) {

	var mountName string
	var mountHeight int
	var temperature int

	fmt.Println("Далее введите данные о восхождение")

	fmt.Println("Как называется ваша гора?")
	fmt.Scan(&mountName)

	fmt.Println("Какая высота у этой горы?")
	fmt.Scan(&mountHeight)

	fmt.Println("Какая температура ожидается во время восхождени?")
	fmt.Scan(&temperature)

	return mountName, mountHeight, temperature
}

func mountainDifficultyLevel(checkAge bool, checkEquipment bool, checkExpirience bool, mountHeight int, temperature int) (bool, string) {

	easyLevel := "Легкий уровень"
	mediumLevel := "Средний уровень"
	hardLevel := "Сложный уровень"

	switch {
	case checkAge && checkEquipment && checkExpirience && mountHeight >= 7000 && temperature <= -30:
		fmt.Println("Вам подходит сложный уровень гор")
		return true, hardLevel
	case checkAge && checkEquipment && mountHeight >= 5000 && temperature <= -15:
		fmt.Println("Вам подходит средний уровень гор")
		return true, mediumLevel
	case checkAge && mountHeight <= 2000 && temperature >= 10:
		fmt.Println("Вам подходит легкий уровень сложности гор")
		return true, easyLevel
	default:
		return false, ""
	}

}

func alpCard(name string, age int, weight float64, height int, expirience int, helmet string, boots string, insurance string, empty bool, backpackWeight float64, difficultyLevel string, mountName string, mountHeight int, temperature int) {

	fmt.Println("----------------------Карточка альпиниста----------------------")
	fmt.Println("Имя: ", name, "\n Возраст: ", age, "\n Ваш вес: ", weight, "\n Ваш рост:", height, "\n Опыт альпинизма: ", expirience, "\n Есть каска: ", helmet, "\n Есть горные ботинки: ", boots, "\n Есть ли страховка: ", insurance, "\n Вес рюкзака: ", backpackWeight, "\n Название горы: ", mountName, "\n Высота Горы: ", mountHeight, "\n Температура во время восхождения: ", temperature, "\n Ваша допустимая сложность восхождения: ", difficultyLevel, "\n Удачного восхождения!!! :)")
}

/* func weatherCheck() {

	var weather string
	var temperature int
	var mount string
	var mountHeight int

	fmt.Println("Какую гору ты хочешь покорить?")
	fmt.Scan(&mount)

	fmt.Println("Какая у нее высота?")
	fmt.Scan(&mountHeight)

	fmt.Println("Какая ожидается погода?")
	fmt.Scan(&weather)

	fmt.Println("Какая ожидается температура?")
	fmt.Scan(&temperature)
}

/*
func main() {

	var name string
	var age int
	var experience int
	var height int
	var weight int
	var hasHelmet string
	var hasBoots string
	var hasInsurance string
	var backpackWeight int
	var temperature int
	var water int

	fmt.Println("Введите ваше имя:")
	fmt.Scan(&name)

	fmt.Println("Введите ваш возраст:")
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Возраст подходит")
	} else {
		fmt.Println("Регистрация невозможно")
		return
	}

	fmt.Println("Сколько лет опыта в альпинизме?")
	fmt.Scan(&experience)

	if experience >= 2 {
		fmt.Println("Достаточно опыта")
	} else {
		fmt.Println("Недостаточно опыта")
		return
	}

	fmt.Println("Ваш рост (см):")
	fmt.Scan(&height)

	fmt.Println("Ваш вес (кг):")
	fmt.Scan(&weight)

	fmt.Println("Есть ли у вас каска? (да/нет)")
	fmt.Scan(&hasHelmet)

	if hasHelmet == "да" {
		fmt.Println("Каска есть")
	} else {
		fmt.Println("Без каски восхождение запрещено")
		return
	}

	fmt.Println("Есть ли у вас ботинки? (да/нет)")
	fmt.Scan(&hasBoots)

	if hasBoots == "да" {
		fmt.Println("Горные ботинки есть")
	} else {
		fmt.Println("Нужны специальные ботинки")
		return
	}

	fmt.Println("Есть ли страховка? (да/нет)")
	fmt.Scan(&hasInsurance)

	if hasInsurance == "да" {
		fmt.Println("Страховка оформлена")
	} else {
		fmt.Println("Нужно оформить страховку")
		return
	}

	fmt.Println("Вес рюкзака (кг):")
	fmt.Scan(&backpackWeight)

	if backpackWeight <= 20 {
		fmt.Println("Вес рюкзака нормальный")
	} else {
		fmt.Println("Рюкзак слишком тяжёлый")
		return
	}

	fmt.Println("Температура на маршруте:")
	fmt.Scan(&temperature)

	if temperature < 0 {
		fmt.Println(" На улице мороз")
	} else {
		fmt.Println(" Погода тёплая")
	}

	fmt.Println("Количество воды (литры):")
	fmt.Scan(&water)

	if water >= 2 {
		fmt.Println(" Воды достаточно")
	} else {
		fmt.Println("Возьмите больше воды")
	}

	fmt.Println("\n--- Карточка альпиниста ---")
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d лет\n", age)
	fmt.Printf("Опыт: %d лет\n", experience)
	fmt.Printf("Рост: %d см\n", height)
	fmt.Printf("Вес: %d кг\n", weight)
	fmt.Printf("Каска: %s\n", hasHelmet)
	fmt.Printf("Ботинки: %s\n", hasBoots)
	fmt.Printf("Страховка: %s\n", hasInsurance)
	fmt.Printf("Вес рюкзака: %d кг\n", backpackWeight)
	fmt.Printf("Температура: %d °C\n", temperature)
	fmt.Printf("Вода: %d литров\n", water)

}
*/
