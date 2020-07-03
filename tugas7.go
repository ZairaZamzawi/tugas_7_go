package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	var number = 10
	var word = "Ini Kata"

	bacatipe1(number)
	go bacatipe1(number)

	bacatipe2(word)
	go bacatipe2(word)

	var input string
	fmt.Scanln(&input)
}

func bacatipe1(n int) {
	var reflectNumber = reflect.ValueOf(n)
	fmt.Println("Tipe Data Pertama", reflectNumber.Type())
}

func bacatipe2(w string) {
	var reflectNumber = reflect.ValueOf(w)
	fmt.Println("Tipe Data Kedua", reflectNumber.Type())
}
