package main

import ()

func changeString(str *string) {
	*str = "Other string"
}

func removeLastChar(str *string) {
	*str = (*str)[:len(*str)-1]
}

func main() {
	str := "Pointer String"
	println(str)
	println(&str)

	changeString(&str)
	println(str)
	removeLastChar(&str)
	println(str)
}
