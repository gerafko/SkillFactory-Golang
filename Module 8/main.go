package main

import (
	"electronic"
	"fmt"
)

func main() {
	Apple := electronic.NewApplePhone("13 Pro Max")
	Android := electronic.NewAndroidPhone("Google", "Pixel 6")
	Station := electronic.NewRadioPhone("Philips", "P350", 13)

	printCharacteristics(Apple)
	printCharacteristics(Android)
	printCharacteristics(Station)
}

func printCharacteristics(in electronic.Phone) {
	switch in.Type() {
	case "station":
		radio := in.(electronic.StationPhone)
		fmt.Printf("Бренд:%s Модель:%s Тип:%s Кол-во кнопок:%s \n", in.Brand(), in.Model(), in.Type(), radio.ButtonsCount())
	case "smartphone":
		smart := in.(electronic.Smartphone)
		fmt.Printf("Бренд:%s Модель:%s Тип:%s Ось:%s \n", in.Brand(), in.Model(), in.Type(), smart.OS())
	}
}
