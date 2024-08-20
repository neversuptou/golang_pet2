package main

import (
	"fmt"
)

var marks = map[string]string{"qq": "bvd"}

func main() {
	for {
		viewMenu()
		menuKey := scanMenu()
		if menuKey <= 0 || menuKey >= 4 {
			break
		}

	}
}

func viewMenu() (key int) {
	fmt.Println("1. Посмотреть закладки")
	fmt.Println("2. Добавить закладку")
	fmt.Println("3. Удалить закладку")
	fmt.Println("4. Выход")
	return
}

func scanMenu() (key int) {
	fmt.Scan(&key)
	switch key {
	case 1:
		lookMarks(marks)
	case 2:
		addMark(marks)
	case 3:
		deleteMark()
	default:
		return
	}
	return
}

func lookMarks(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
func addMark(m map[string]string) {
	var newKey string
	var newDisc string
	fmt.Println("Введите название закладки")
	fmt.Scan(&newKey)
	if !checkDuplicate(newKey, marks) {
		fmt.Println("Введите описание")
		fmt.Scan(&newDisc)
		m[newKey] = newDisc
		fmt.Printf("Закладка %v успещно создана \n", newKey)
	} else {
		fmt.Println("Закладка уже создана")
	}
}

func checkDuplicate(newKey string, m map[string]string) bool {
	var check bool
	for key, _ := range m {
		if key == newKey {
			check = true
		} else {
			check = false
		}
	}
	return check
}

func deleteMark() {
	var newKey string
	fmt.Println("Введите название закладки для удаления")
	fmt.Scan(&newKey)
	if checkDuplicate(newKey, marks) {
		delete(marks, newKey)
		fmt.Printf("Закладка %v успешно удалена \n", newKey)
	} else {
		fmt.Println("Такой закладки не существует, проверьте список закладок клавишой 1")
	}
}
