package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Name       string
	Age        int
	Weight     float64
	Height     int
	Experience int
}

type Equipment struct {
	Helmet    bool
	Boots     bool
	Insurance bool
	BpWeight  float64
}

type MountInfo struct {
	MountName   string
	MountHeight int
	Temperature int
}

func main() {

	fmt.Println("Привет альпинист получи свою карточку для покорения вершин  :) ")
	fmt.Println("Для начала расскажи мне о себе")

	u := User{}
	e := Equipment{}
	mi := MountInfo{}

	err := InputUser(&u, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = checkWeight(&u, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	inputMountInfo(&mi)
}

func InputUser(u *User, e *Equipment) error {

	var input string

	fmt.Println("Как тебя зовут?")
	fmt.Scan(&u.Name)

	fmt.Println("Сколько тебе лет?")
	fmt.Scan(&u.Age)

	err := checkAge(u)
	if err != nil {
		return err
	}

	fmt.Println("Какой у тебя вес?")
	fmt.Scan(&u.Weight)

	fmt.Println("Какой у тебя рост?")
	fmt.Scan(&u.Height)

	fmt.Println("Какой у тебя опыт покорений гор в годах?")
	fmt.Scan(&u.Experience)

	fmt.Println("Есть ли у тебя каска? (да/нет)")
	fmt.Scan(&input)

	helmet, err := parseAnswer(input)
	if err != nil {
		return err
	}
	e.Helmet = helmet

	fmt.Println("Есть ли у тебя горные ботинки? (да/нет)")
	fmt.Scan(&input)

	boots, err := parseAnswer(input)
	if err != nil {
		return err
	}
	e.Boots = boots

	fmt.Println("Есть ли у тебя страховка? (да/нет)")
	fmt.Scan(&input)

	insurance, err := parseAnswer(input)
	if err != nil {
		return err
	}
	e.Insurance = insurance

	fmt.Println("какой вес у твоего рюказака в кг?")
	fmt.Scan(&e.BpWeight)

	return nil
}

func parseAnswer(input string) (bool, error) {

	switch {
	case strings.EqualFold(input, "Да"):
		return true, nil
	case strings.EqualFold(input, "Нет"):
		return false, nil
	default:
		return false, errors.New("Некорректный ввод")

	}

}

func checkAge(u *User) error {

	if u.Age >= 18 {
	} else {
		return errors.New("Людям младше 18 лет допуск запрещен")
	}
	return nil

}

func checkWeight(u *User, e *Equipment) error {

	var sum float64

	sum = u.Weight * 0.25

	if e.BpWeight > sum {
		return errors.New("Рюкзак слишком тяжелый для тебя")
	}
	return nil
}

func inputMountInfo(mi *MountInfo) {

	fmt.Println("Далее введите данные о предпочтительном восхождении")

	fmt.Println("Введите название горы:")
	fmt.Scan(&mi.MountName)

	fmt.Println("Какая высота у этой горы?")
	fmt.Scan(&mi.MountHeight)

	fmt.Println("Какая температура ожидается во время восхождени?")
	fmt.Scan(&mi.Temperature)
}

func mountainDifficultyLevel(checkAge bool /*checkEquipment bool,*/, expirience int, mountHeight int, temperature int, helmet string, boots string, insurance string) (bool, string) {

	easyLevel := "Легкий уровень"
	mediumLevel := "Средний уровень"
	hardLevel := "Сложный уровень"

	switch {
	case checkAge && helmet == "да" && boots == "да" && insurance == "да" && expirience >= 3 && mountHeight >= 7000 && temperature <= -30:
		fmt.Println("Вам подходит сложный уровень гор")
		return true, hardLevel
	case checkAge && expirience >= 2 && mountHeight <= 6999 && temperature >= -29 && ((helmet == "да" && insurance == "да") || (boots == "да" && insurance == "да")):
		fmt.Println("Вам подходит средний уровень сложности")
		return true, mediumLevel
	case checkAge && expirience >= 1 && mountHeight <= 2000 && temperature >= 0 && (helmet == "да" || boots == "да" || insurance == "да"):
		fmt.Println("Вам подходит легкий уровень сложности гор")
		return true, easyLevel
	default:
		fmt.Println("Человек написавший этот код тупой и не учел сложившуюся совокупность факторов и ему было лень ее переделывать так что вы не попадаете не под какие уровни")
		return true, "зато вы посмотрели карточку вау круто!!! во второй версии обещаю переделать :))"
	}
}
func alpCard(name string, age int, weight float64, height int, expirience int, helmet string, boots string, insurance string, backpackWeight float64, difficultyLevel string, mountName string, mountHeight int, temperature int) {

	fmt.Println("----------------------Карточка альпиниста----------------------")
	fmt.Println("Имя: ", name, "\n Возраст: ", age, "\n Ваш вес: ", weight, "\n Ваш рост:", height, "\n Опыт альпинизма: ", expirience, "\n Есть каска: ", helmet, "\n Есть горные ботинки: ", boots, "\n Есть ли страховка: ", insurance, "\n Вес рюкзака: ", backpackWeight, "\n Название горы: ", mountName, "\n Высота Горы: ", mountHeight, "\n Температура во время восхождения: ", temperature, "\n Ваша допустимая сложность восхождения: ", difficultyLevel, "\n Удачного восхождения!!! :)")
}
