package main
import "fmt"
func pointer(){
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println("numberA (value)   :", numberA)  // 4
	fmt.Println("numberA (address) :", &numberA) // 0xc20800a220
	fmt.Println("numberB (value)   :", *numberB) // 4
	fmt.Println("numberB (address) :", numberB)  // 0xc20800a220
	fmt.Println("")
	numberA = 5
	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	var number = 4
	fmt.Println("before",number)
	change(&number, 10)
	fmt.Println("after",number)
}

func change(original *int, value int){
	*original = value
}