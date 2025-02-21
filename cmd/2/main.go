package main

import "fmt"

func main() {
	var intNum int = 2623

	intNum = intNum + 1

	fmt.Println(intNum)

	var floatNum float32 = 1212.42

	fmt.Println(floatNum)

	var myString string = " hello \n" + " " + "World"

	fmt.Println(myString)

	var myRune rune = 'a'

	fmt.Println(myRune)

	myFancy := "hey"

	fmt.Println(myFancy)

	var1, var2 := 1, 2

	fmt.Println(var1, var2)

	const myConst string = "3.14"

	fmt.Println(myConst)
}
