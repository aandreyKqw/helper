package main

import (
	"fmt"
	"os"
)

func main() {
	open_service()
}

func open_service() {
	if len(os.Args) == 1 {
		fmt.Println("нет аргументов")
		return
	}

	switch os.Args[1] {
	case "--write":
		writeData()
	case "--static":
		// Защита от паники, если второй аргумент не передан
		if len(os.Args) < 3 {
			fmt.Println("no value")
			return
		}
		switch os.Args[2] {
		case "7days":
			showData("day")
		case "mounth":
			showData("mnt")
		default:
			fmt.Println("no value")
			return
		}
	default:
		fmt.Println("аргументы только --write or --static")
	}
}

func writeData() {
	fmt.Println("Записываем в БД")
}

func showData(value string) {
	switch value {
	case "day":
		fmt.Println("Вывод показателей за неделю")
	case "mnt":
		fmt.Println("Вывод показателей за месяц")
	}
}
